package rmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
}

func NewRabbitMQ(cfg Config) (*amqp.Connection, error) {
	return amqp.Dial(getURL(cfg))
}

func getURL(cfg Config) string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s/",
		"test", "test", "localhost:5672",
	)
}
