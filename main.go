package main

import (
	"context"
	"encoding/json"
	"fmt"
	"rabitmq/config"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// "github.com/aymanbagabas/go-udiff"
	// a := "Hello, world!\n"
	// b := "Hello, world! Say hi to µDiff\n"
	// d := udiff.Unified("a.txt", "b.txt", a, b)
	// fmt.Println(d)

	// rabit mq server
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
	// body := "hello aftab!"
	h := config.Human{Name: "hari", Age: 13}
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
	if err := configr(); err != nil {
		fmt.Print("helllo")
	}
}
