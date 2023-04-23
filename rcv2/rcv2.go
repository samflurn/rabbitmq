package main

import (
	"log"
	"rabitmq/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
func main() {
	conn, err := amqp.Dial("amqp://127.0.0.1")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	ch.ExchangeDeclare(string(config.FanOutEXchange), "fanout", true, false, false, false, nil)
	q, err := ch.QueueDeclare(
		config.AnotherQueue2, // name
		true,                 // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	failOnError(err, "Failed to declare a queue")
	ch.QueueBind(q.Name, config.HelloQueueRoute, string(config.FanOutEXchange), false, nil)
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			// payload := &config.Human{}
			// if err := json.Unmarshal(d.Body, payload); err != nil {
			// 	log.Print(err)
			// }

			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
