package main

import (
	"log"

	"github.com/slonegd-otus-go/12_calendar/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
