package amqpsubscriber

import (
	"log"

	"github.com/streadway/amqp"
)

func Run(url, name string, onEvent func(string)) {
	log.Printf("start event subscriber on %s", url)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Printf("amqp dial failed: %s", err)
		return
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Printf("open amqp channel failed: %s", err)
		return
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Printf("declare queue failed: %s", err)
		return
	}

	delivery, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Printf("register a amqp consumer failed: %s", err)
		return
	}

	for message := range delivery {
		onEvent(string(message.Body))
	}
}
