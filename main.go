package main

import (
	"log"
	"net"
	"os"
	"strconv"

	"github.com/slonegd-otus-go/12_calendar/internal/event"
	proto "github.com/slonegd-otus-go/12_calendar/internal/grpc"
	"github.com/slonegd-otus-go/12_calendar/internal/web"
	"google.golang.org/grpc"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("PORT not defined")
	}

	storage := event.NewStorage()

	go web.Run(port, storage)

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterCalendarServer(grpcServer, proto.NewServer(storage))
	grpcServer.Serve(listener)
}
