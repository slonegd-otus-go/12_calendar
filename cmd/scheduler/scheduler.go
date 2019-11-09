package scheduler

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/slonegd-otus-go/12_calendar/internal/event"
	"github.com/slonegd-otus-go/12_calendar/internal/event/amqppublisher"
	"github.com/slonegd-otus-go/12_calendar/internal/event/psqlstorage"
)

var ampqURL string

func init() {
	Command.Flags().StringVar(&ampqURL, "ampqurl", "amqp://guest:guest@localhost:5672", "url to ampq server")
}

var Command = &cobra.Command{
	Use:   "scheduler",
	Short: "Run event scheduler (amqp publisher)",
	Run: func(cmd *cobra.Command, args []string) {
		storage := psqlstorage.New()
		publisher := amqppublisher.New(ampqURL)
		log.Printf("start event ampq publisher on %s", ampqURL)
		event.Scan(storage, publisher.OnEvent)
	},
}
