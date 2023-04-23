package config

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type (
	Human struct {
		Name string
		Age  int
	}
	ExchangeName string
	QueueName    string
)

const (
	FanOutEXchange  ExchangeName = "fan-out-exchange"
	HelloQueue                   = "hello"
	HelloQueueRoute              = "hello-route"
	AnotherQueue                 = "other-queue"
	AnotherQueue2                = "other-queue2"
)

func Init() (*amqp.Channel, *amqp.Connection, error) {

	conn, err := amqp.Dial("amqp://127.0.0.1")
	if err != nil {
		fmt.Print(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		fmt.Print(err)
	}
	return ch, conn, nil
}
