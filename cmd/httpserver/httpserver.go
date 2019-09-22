package httpserver

import (
	"github.com/spf13/cobra"

	"github.com/slonegd-otus-go/12_calendar/internal/event/mapstorage"
	"github.com/slonegd-otus-go/12_calendar/internal/web"
)

var host string
var port int

var Command = &cobra.Command{
	Use:   "httpserver",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		storage := mapstorage.New()
		web.Run(host, port, storage)
	},
}

func init() {
	Command.Flags().StringVar(&host, "host", "localhost", "host to listen")
	Command.Flags().IntVar(&port, "port", 8080, "port to listen")
}
