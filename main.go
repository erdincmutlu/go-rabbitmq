package main

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	// amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Printf("Go RabbitMQ\n")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to Rabbit MQ instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("TestQueue", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf("%+v\n", q)

	err = ch.PublishWithContext(context.TODO(), "", "TestQueue", false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Successfully published message to queue")

}
