package config

import "fmt"

// RabbitmqURL returns url for RabbitMQ connection
func RabbitmqURL() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s/",
		"test", "test", "localhost:5672",
	)
}

// MongoURL returns url for MongoDB connection
func MongoURL() string {
	return fmt.Sprintf("mongodb://127.0.0.1:27017")
}
