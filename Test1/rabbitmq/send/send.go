package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"strconv"
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

	queue, err := ch.QueueDeclare("first hello world", false, false, false, false, nil)
	if err != nil {
		println("queue err:" + err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body0 := "hello world"
	var body string
	for i := 0; i < 10; i++ { //发送十次
		body = body0 + strconv.Itoa(i)
		err = ch.PublishWithContext(ctx,
			"",         // exchange
			queue.Name, // routing key
			false,      // mandatory
			false,      // immediate
			amqp.Publishing{
				//DeliveryMode: amqp.Persistent, //2为持久化到磁盘 0或1为不持久化 默认为0
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		if err != nil {
			println("publish err:" + err.Error())
			return
		}
	}

}
