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
	defer ch.Close()
	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		println("err:" + err.Error())
		return
	}
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		println("err:" + err.Error())
		return
	}
	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil,
	)
	if err != nil {
		println("err:" + err.Error())
		return
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		println("err:" + err.Error())
		return
	}
	var forever chan struct{}

	//go func() {
	for d := range msgs {
		log.Printf(" [x] %s", d.Body)
	}
	//}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
