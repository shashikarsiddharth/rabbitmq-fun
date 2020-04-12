package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("message published to queue!")

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

	// publish message
	// body := []byte("Hello World")
	body, err := readFile()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan error)
	for i := 0; i < 1000; i++ {
		go publish(i, ch, queue, body, done)
	}

	fmt.Println(<-done)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func readFile() ([]byte, error) {
	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, err
	}
	return data, nil
}

func publish(ix int, ch *amqp.Channel, q amqp.Queue, data []byte, done chan<- error) {
	t := time.NewTicker(time.Duration(1) * time.Second)
	for {
		<-t.C
		fmt.Printf("%d is publishing\n", ix)
		err := ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        data,
			})
		if err != nil {
			done <- err
		}
	}
}
