package rabbitmq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Sendmsg(body string) {
	err := ch.PublishWithContext(
		context.Background(),
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		panic(err)
	}
	sentCount++
}

func Consumemsg() <-chan amqp.Delivery {
	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		panic(err)
	}
	return msgs
}
