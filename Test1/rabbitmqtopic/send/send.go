package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func main() {
	conn, err := amqp.Dial("amqp://162.14.64.254:5672/")
	if err != nil {
		println("conn err:" + err.Error())
		return
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		println("err:" + err.Error())
		return
	}
	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		println("err:" + err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body := "你好世界"
	err = ch.PublishWithContext(ctx,
		"logs_topic", // exchange
		"wei.ren",    // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
}
