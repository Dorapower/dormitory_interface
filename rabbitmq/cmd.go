package rabbitmq

import (
	"context"
	"dormitory_interface/sql"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"strings"
	"time"
)

var ch *amqp.Channel
var queue amqp.Queue
var sentCount, receivedCount int

func init() {
	conn, err := amqp.Dial("amqp://user01:123456@47.92.123.159:5672/")
	if err != nil {
		panic(err)
	}
	ch, err = conn.Channel()
	if err != nil {
		panic(err)
	}
	queue, err = ch.QueueDeclare(
		"order",      // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		amqp.Table{}, // arguments
	)
	if err != nil {
		panic(err)
	}
	err = ch.PublishWithContext(
		context.Background(),
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("hello"),
		})
	if err != nil {
		panic(err)
	}
	log.Println("rabbitmq init success")
	go func() {
		for d := range Consumemsg() {
			//split on space
			data := string(d.Body)
			if !strings.Contains(data, " ") {
				continue
			}
			username, buildingNo := strings.Split(string(d.Body), " ")[0], strings.Split(string(d.Body), " ")[1]
			sql.InsertApplicationRaw(username, buildingNo)
			receivedCount++
		}
	}()
	go func() {
		for {
			log.Printf("Sent: %d, Received: %d", sentCount, receivedCount)
			sentCount, receivedCount = 0, 0
			time.Sleep(1 * time.Second)
		}
	}()
}
