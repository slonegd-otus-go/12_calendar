package api

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/slonegd-otus-go/12_calendar/internal/event"
	"github.com/slonegd-otus-go/12_calendar/internal/event/mapstorage"
	"github.com/slonegd-otus-go/12_calendar/internal/event/psqlstorage"
	proto "github.com/slonegd-otus-go/12_calendar/internal/grpc"
	"github.com/slonegd-otus-go/12_calendar/internal/web"
)

var host string
var port int
var isGRPC bool
var storageType string
var connection string

func init() {
	Command.Flags().StringVar(&host, "host", "127.0.0.1", "host to listen")
	Command.Flags().IntVar(&port, "port", 8080, "port to listen")
	Command.Flags().BoolVar(&isGRPC, "grpc", false, "set for grpc api instead rest api")
	Command.Flags().StringVar(&storageType, "storage", "psql", "storage type (map or psql)")
	Command.Flags().StringVar(&connection, "connection", "host=localhost user=myuser password=mypass dbname=mydb", "connection string for postgresql")
}

var Command = &cobra.Command{
	Use:   "api",
	Short: "Run api server",
	Run: func(cmd *cobra.Command, args []string) {

		var storage event.Storage
		switch storageType {
		case "map":
			storage = mapstorage.New()
		case "psql":
			storage = psqlstorage.New(connection)
		default:
			log.Fatalf("unknow storage type, want map or psql, got %s", storageType)
		}

		if isGRPC {
			address := fmt.Sprintf("%s:%d", host, port)
			listener, err := net.Listen("tcp", address)
			if err != nil {
				log.Fatalf("failed to listen %v", err)
			}

			grpcServer := grpc.NewServer()
			proto.RegisterCalendarServer(grpcServer, proto.NewServer(storage))
			log.Printf("start grpc server")
			grpcServer.Serve(listener)
			return
		}

		log.Printf("start rest api server")
		web.Run(host, port, storage)

	},
}
