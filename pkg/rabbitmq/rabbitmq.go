package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (channel *amqp.Channel, err error) {
	conn, err := amqp.Dial("amqp://user:pass@localhost:5672/")
	if err != nil {
		panic(err)
	}

	channel, err = conn.Channel()
	if err != nil {
		panic(err)
	}

	return channel, nil
}

func Consume(channel *amqp.Channel, out chan<- amqp.Delivery) error {
	msgs, err := channel.Consume(
		"test-queue",    // queue
		"test-consumer", // consumer
		false,           // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		panic(err)
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}
