package main

import (
	"encoding/json"
	"faker/internal/action"
	"fmt"
	"log"
	"os"
	"time"

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
	duration := action.GetDuration(os.Args)
	log.Println("Faker started with interval: ", duration)

	url := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		viper.GetString("amqp.user"),
		viper.GetString("amqp.password"),
		viper.GetString("amqp.host"),
		viper.GetString("amqp.port"),
	)
	conn, err := amqp.Dial(url)
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
		log.Fatal("Failed to declare an Alarm queue ", err)
	}

	for {
		alarm := action.GetRandomAlarm()
		body, err := json.Marshal(alarm)
		if err != nil {
			log.Fatal("Failed to marshal alarm ", err)
		}
		err = publish(ch, q.Name, body)
		if err != nil {
			log.Fatal("Failed to publish a message ", err)
		}

		log.Printf("Sent to %s queue: %s", q.Name, body)
		time.Sleep(duration * time.Second)
	}
}

func publish(ch *amqp.Channel, queueName string, body []byte) error {
	return ch.Publish("", queueName, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
}
