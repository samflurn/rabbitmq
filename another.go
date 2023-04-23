package main

import (
	"context"
	"encoding/json"
	"fmt"
	"rabitmq/config"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func configr() error {
	ch, conn, err := config.Init()
	defer ch.Close()
	defer conn.Close()
	q, err := ch.QueueDeclare(
		config.HelloQueue, // name
		false,             // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// body := "hello sam!"
	h := config.Human{Name: "sam", Age: 12}
	payload, _ := json.Marshal(h)
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        payload,
		})
	if err != nil {
		fmt.Print(err)
	}
	return err
}
