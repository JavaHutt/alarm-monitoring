package main

import (
	"context"
	"log"
	"monitor/internal/adaptor"
	"monitor/internal/rmq"
	"monitor/internal/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Failed to initialize config: %s", err.Error())
	}

	ctx := context.Background()

	db, err := adaptor.NewMongoDB(ctx, adaptor.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Client().Disconnect(ctx)

	adaptor := adaptor.NewAdaptor(db)
	service := service.NewService(adaptor)

	conn, err := rmq.NewRabbitMQ(rmq.Config{
		Host:     viper.GetString("amqp.host"),
		Port:     viper.GetString("amqp.port"),
		User:     viper.GetString("amqp.user"),
		Password: viper.GetString("amqp.password"),
	})
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel ", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("Alarm", false, false, false, false, nil)
	if err != nil {
		log.Fatal("Failed to declare a queue ", err)
	}

	msgs, err := ch.Consume(q.Name, "consumer", false, false, false, false, nil)
	if err != nil {
		log.Fatal("Failed to register a consumer ", err)
	}

	forever := make(chan os.Signal)
	signal.Notify(forever, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool)

	go recieveAlarms(ctx, service, msgs, done)

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever

	ch.Cancel("consumer", false)
	log.Println("Stopped receiving message from queue")
	log.Println("Wait for worker process recieved message")
	<-done
	log.Println("Woker done")
}

func recieveAlarms(ctx context.Context, service *service.Service, msgs <-chan amqp.Delivery, done chan bool) {
	for m := range msgs {
		parsed, err := service.ParseAlarm(m.Body)
		if err != nil {
			log.Fatal("Failed parse incoming alarm ", err)
		}

		err = service.InsertAlarm(ctx, *parsed)
		if err != nil {
			log.Fatal("Failed insert incoming alarm ", err)
		}

		m.Ack(false)
		log.Println("Alarm acked")
	}
	done <- true
}
