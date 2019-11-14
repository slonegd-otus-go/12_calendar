package sender

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/slonegd-otus-go/12_calendar/internal/event/amqpsubscriber"
)

var amqpURL string

func init() {
	Command.Flags().StringVar(&amqpURL, "amqpurl", "amqp://guest:guest@localhost:5672", "url to amqp server")
}

var Command = &cobra.Command{
	Use:   "sender",
	Short: "Run event sender (amqp subscriber)",
	Run: func(cmd *cobra.Command, args []string) {
		amqpsubscriber.Run(amqpURL, "event", func(message string) {
			log.Printf("got event: %s", message)
		})
	},
}
