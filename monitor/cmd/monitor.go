package main

import (
	"context"
	"fmt"
	"log"
	"monitor/internal/adaptor"
	"time"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Failed to initialize config: %s", err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := adaptor.NewMongoDB(ctx, adaptor.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func recieve(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		d.Ack(false)
	}
}
