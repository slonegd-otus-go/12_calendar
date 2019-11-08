package grpcserver

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/slonegd-otus-go/12_calendar/internal/event"
	"github.com/slonegd-otus-go/12_calendar/internal/event/amqppublisher"
	"github.com/slonegd-otus-go/12_calendar/internal/event/mapstorage"
	proto "github.com/slonegd-otus-go/12_calendar/internal/grpc"
)

var host string
var port int
var ampqURL string

func init() {
	Command.Flags().StringVar(&host, "host", "localhost", "host to listen")
	Command.Flags().IntVar(&port, "port", 50051, "port to listen")
	Command.Flags().StringVar(&ampqURL, "ampqurl", "amqp://guest:guest@localhost:5672/", "url to ampq server")
}

var Command = &cobra.Command{
	Use:   "grpcserver",
	Short: "Run grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		address := fmt.Sprintf("%s:%d", host, port)
		listener, err := net.Listen("tcp", address)
		if err != nil {
			log.Fatalf("failed to listen %v", err)
		}

		storage := mapstorage.New()

		publisher := amqppublisher.New(ampqURL)
		event.StartScan(storage, publisher.OnEvent)

		grpcServer := grpc.NewServer()
		proto.RegisterCalendarServer(grpcServer, proto.NewServer(storage))
		grpcServer.Serve(listener)
	},
}
