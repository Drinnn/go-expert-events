package main

import (
	"fmt"
	"github.com/Drinnn/go-expert-events/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitMqChannel, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer rabbitMqChannel.Close()

	msgsChannel := make(chan amqp.Delivery)

	go rabbitmq.Consume(rabbitMqChannel, msgsChannel)

	for msg := range msgsChannel {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
