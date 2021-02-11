package config

import "fmt"

// RabbitmqURL return url for RabbitMQ connection
func RabbitmqURL() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s/",
		"test", "test", "localhost:5672",
	)
}
