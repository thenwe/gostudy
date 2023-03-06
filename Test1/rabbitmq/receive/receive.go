package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
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
	queue, err := ch.QueueDeclare(
		"first hello world",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		println("queue err:" + err.Error())
		return
	}
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
		println("queue err:" + err.Error())
		return
	}
	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	<-forever //一直阻塞
}
