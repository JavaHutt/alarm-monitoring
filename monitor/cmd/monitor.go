package main

import (
	"log"
	"monitor/config"

	"github.com/streadway/amqp"
)

func main() {
	rmqURL := config.RabbitmqURL()
	conn, err := amqp.Dial(rmqURL)
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

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal("Failed to register a consumer ", err)
	}

	forever := make(chan bool)

	go recieve(msgs)

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
}

func recieve(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		d.Ack(false)
	}
}
