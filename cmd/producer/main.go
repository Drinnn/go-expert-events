package main

import (
	"fmt"
	"github.com/Drinnn/go-expert-events/pkg/rabbitmq"
)

func main() {
	rabbitMQChannel, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer rabbitMQChannel.Close()

	// get input from user and publish it
	for {
		var message string
		fmt.Scanln(&message)
		rabbitmq.Publish(rabbitMQChannel, []byte(message))
	}
}
