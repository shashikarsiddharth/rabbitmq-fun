package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Hello World!")

	// connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@0.0.0.0:5672")
	failOnError(err, "Failed to connect to RabbitMQ! ")
	defer conn.Close()

	// open channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open channel! ")
	defer ch.Close()

	// declare queue
	queue, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue! ")

	// consume messages
	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		true,
		true,
		true,
		nil,
	)
	failOnError(err, "Failed to declare consumer! ")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
