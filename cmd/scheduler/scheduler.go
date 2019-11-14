package scheduler

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/slonegd-otus-go/12_calendar/internal/event"
	"github.com/slonegd-otus-go/12_calendar/internal/event/amqppublisher"
	"github.com/slonegd-otus-go/12_calendar/internal/event/psqlstorage"
)

var amqpURL string
var connection string

func init() {
	Command.Flags().StringVar(&amqpURL, "amqpurl", "amqp://guest:guest@localhost:5672", "url to ampq server")
	Command.Flags().StringVar(&connection, "connection", "host=localhost user=myuser password=mypass dbname=mydb", "connection string for postgresql")
}

var Command = &cobra.Command{
	Use:   "scheduler",
	Short: "Run event scheduler (amqp publisher)",
	Run: func(cmd *cobra.Command, args []string) {
		storage := psqlstorage.New(connection)
		publisher := amqppublisher.New(amqpURL, "event")
		log.Printf("start event ampq publisher on %s", amqpURL)
		event.Scan(storage, publisher.OnEvent)
	},
}
