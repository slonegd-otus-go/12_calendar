package cmd

import (
	"github.com/spf13/cobra"

	"github.com/slonegd-otus-go/12_calendar/cmd/grpcclient"
	"github.com/slonegd-otus-go/12_calendar/cmd/grpcserver"
	"github.com/slonegd-otus-go/12_calendar/cmd/httpserver"
)

var Command = &cobra.Command{
	Use:   "mycalendar",
	Short: "calendar is a calendar micorservice demo",
}

func init() {
	Command.AddCommand(httpserver.Command)
	Command.AddCommand(grpcserver.Command)
	Command.AddCommand(grpcclient.Command)
}
