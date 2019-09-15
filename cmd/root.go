package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "calendar",
	Short: "calendar is a calendar micorservice demo",
}

func init() {
	RootCmd.AddCommand(HTTPserverCommand)
}
