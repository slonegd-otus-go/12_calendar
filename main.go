package main

import (
	"log"
	"os"
	"strconv"

	"github.com/slonegd-otus-go/12_calendar/internal/event"
	"github.com/slonegd-otus-go/12_calendar/internal/web"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("PORT not defined")
	}

	storage := event.NewStorage()

	web.Run(port, storage)
}
