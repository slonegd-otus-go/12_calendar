package amqppublisher

import (
	"log"

	"github.com/streadway/amqp"

	"github.com/slonegd-otus-go/12_calendar/internal/event"
)

type Publisher struct {
	url string
}

func New(url string) *Publisher {
	return &Publisher{url}
}

func (publisher Publisher) OnEvent(event event.Event) {
	log.Printf("start publish event: %v", event)
	conn, err := amqp.Dial(publisher.url)
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
		"event", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Printf("declare queue failed: %s", err)
		return
	}

	err = channel.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event.Description),
		},
	)
	if err != nil {
		log.Printf("amqp publish failed: %s", err)
		return
	}
	log.Printf("publish event success")
}
