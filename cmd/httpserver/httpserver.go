package httpserver

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/slonegd-otus-go/12_calendar/internal/event"
	"github.com/slonegd-otus-go/12_calendar/internal/event/amqppublisher"
	"github.com/slonegd-otus-go/12_calendar/internal/event/mapstorage"
	"github.com/slonegd-otus-go/12_calendar/internal/event/psqlstorage"
	"github.com/slonegd-otus-go/12_calendar/internal/web"
)

var host string
var port int
var storageType string
var ampqURL string

func init() {
	Command.Flags().StringVar(&host, "host", "localhost", "host to listen")
	Command.Flags().IntVar(&port, "port", 8080, "port to listen")
	Command.Flags().StringVar(&storageType, "storage", "map", "storage type (map or psql)")
	Command.Flags().StringVar(&ampqURL, "ampqurl", "amqp://guest:guest@localhost:5672/", "url to ampq server")
}

var Command = &cobra.Command{
	Use:   "httpserver",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		var storage event.Storage
		switch storageType {
		case "map":
			storage = mapstorage.New()
		case "psql":
			storage = psqlstorage.New()
		default:
			log.Fatalf("unknow storage type, want map or psql, got %s", storageType)
		}
		publisher := amqppublisher.New(ampqURL)
		event.StartScan(storage, publisher.OnEvent)
		web.Run(host, port, storage)
	},
}
