package main

import (
	"encoding/json"
	"faker/config"
	"faker/internal/action"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// duration := action.GetDuration(os.Args)

	alarm := action.GetRandomAlarm()
	body, err := json.Marshal(alarm)
	if err != nil {
		log.Fatal("Failed to marshal alarm ", err)
	}

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
		log.Fatal("Failed to declare an Alarm queue ", err)
	}

	err = publish(ch, q.Name, body)
	if err != nil {
		log.Fatal("Failed to publish a message ", err)
	}

	log.Printf("Sent to %s queue: %s", q.Name, body)
}

func publish(ch *amqp.Channel, queueName string, body []byte) error {
	return ch.Publish("", queueName, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
}
