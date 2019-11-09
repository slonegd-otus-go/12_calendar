package cmd

import (
	"github.com/spf13/cobra"

	"github.com/slonegd-otus-go/12_calendar/cmd/api"
	"github.com/slonegd-otus-go/12_calendar/cmd/grpcclient"
	"github.com/slonegd-otus-go/12_calendar/cmd/scheduler"
	"github.com/slonegd-otus-go/12_calendar/cmd/sender"
)

var Command = &cobra.Command{
	Use:   "mycalendar",
	Short: "calendar is a calendar micorservice demo",
}

func init() {
	Command.AddCommand(api.Command)
	Command.AddCommand(scheduler.Command)
	Command.AddCommand(sender.Command)

	// for debug
	Command.AddCommand(grpcclient.Command)
}
