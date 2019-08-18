package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/slonegd-otus-go/12_calendar/web"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("PORT not defined")
	}

	http.HandleFunc("/", web.Handler)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
