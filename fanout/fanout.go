package main

import (
	"context"
	"fmt"
	"rabitmq/config"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, conn, err := config.Init()
	if err != nil {
		fmt.Print("hello errord")
	}
	if err := ch.ExchangeDeclare(
		string(config.FanOutEXchange),
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		fmt.Print("Errored out while exchange declaration")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer ch.Close()
	defer conn.Close()
	if err := ch.PublishWithContext(ctx,
		string(config.FanOutEXchange),
		config.HelloQueueRoute,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello main hoon sam"),
		}); err != nil {
		fmt.Print(err)
	}
}
