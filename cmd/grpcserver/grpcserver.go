package grpcserver

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/slonegd-otus-go/12_calendar/internal/event"
	proto "github.com/slonegd-otus-go/12_calendar/internal/grpc"
)

var host string
var port int

func init() {
	Command.Flags().StringVar(&host, "host", "localhost", "host to listen")
	Command.Flags().IntVar(&port, "port", 50051, "port to listen")
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

		storage := event.NewStorage()

		grpcServer := grpc.NewServer()
		proto.RegisterCalendarServer(grpcServer, proto.NewServer(storage))
		grpcServer.Serve(listener)
	},
}
