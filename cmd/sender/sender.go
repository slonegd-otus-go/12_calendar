package sender

import (
	"github.com/spf13/cobra"

	"github.com/slonegd-otus-go/12_calendar/internal/event/amqpsubscriber"
)

var ampqURL string

func init() {
	Command.Flags().StringVar(&ampqURL, "ampqurl", "amqp://guest:guest@localhost:5672", "url to ampq server")
}

var Command = &cobra.Command{
	Use:   "sender",
	Short: "Run event sender (amqp subscriber)",
	Run: func(cmd *cobra.Command, args []string) {
		amqpsubscriber.Run(ampqURL)
	},
}
