package cmd

import (
	"github.com/spf13/cobra"

	"github.com/slonegd-otus-go/12_calendar/internal/web"
	"github.com/slonegd-otus-go/12_calendar/internal/event"
)

var host string
var port int

var HTTPserverCommand = &cobra.Command{
	Use:   "httpserver",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		storage := event.NewStorage()
	    web.Run(host, port, storage)
	},
}

func init() {
	HTTPserverCommand.Flags().StringVar(&host, "host", "localhost", "host to listen")
	HTTPserverCommand.Flags().IntVar(&port, "port", 8080, "port to listen")
}