//go:generate swagger generate server --target=./internal/web --spec=./api/swagger.yml --exclude-main
//go:generate ffjson ./internal/web/models/event.go
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
