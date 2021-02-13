package main

import (
	"context"
	"log"
	"monitor/internal/adaptor"
	"monitor/internal/model"
	"monitor/internal/service"
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

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	db, err := adaptor.NewMongoDB(ctx, adaptor.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Client().Disconnect(ctx)

	adaptor := adaptor.NewAdaptor(db)
	services := service.NewService(adaptor)
	a := model.Alarm{}
	services.InsertAlarm(ctx, a)

}

func recieve(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		d.Ack(false)
	}
}
