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
		println("ch err:" + err.Error())
		return
	}
	defer ch.Close()
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
		println("ex err:" + err.Error())
		return
	}
	/*q, err := ch.QueueDeclare(
		"newTopic", // name
		false,      // durable
		false,      // delete when unused
		true,       // exclusive
		false,      // no-wait
		nil,        // arguments

	)
	if err != nil {
		println("q err:" + err.Error())
		return
	}*/
	err = ch.QueueBind(
		"newTopic",   // queue name
		"wei.*",      // routing key
		"logs_topic", // exchange
		false,
		nil,
	)
	msgs, err := ch.Consume(
		"newTopic", // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		println("msg err:" + err.Error())
		return
	}
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}
