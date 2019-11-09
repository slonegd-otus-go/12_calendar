package grpcclient

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	proto "github.com/slonegd-otus-go/12_calendar/internal/grpc"
)

var host string
var port int
var command string
var date string
var duration int
var description string
var id int

func init() {
	Command.Flags().StringVar(&host, "host", "localhost", "host to connect")
	Command.Flags().IntVar(&port, "port", 8080, "port to connect")
	Command.Flags().StringVar(&command, "command", "", "command to server create/remove/getlist (required)")
	Command.Flags().StringVar(&date, "date", "2019-09-07 08:20:00", "event start date")
	Command.Flags().IntVar(&duration, "duration", 600, "event duration in seconds")
	Command.Flags().StringVar(&description, "description", "my bithday", "event description")
	Command.Flags().IntVar(&id, "id", 600, "event id, required for remove")
	Command.MarkFlagRequired("command")
}

var Command = &cobra.Command{
	Use:   "grpcclient",
	Short: "Run grpc client",
	Run: func(cmd *cobra.Command, args []string) {
		address := fmt.Sprintf("%s:%d", host, port)
		connection, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("could not connect: %v", err)
		}
		log.Printf("connect to grpc server: %s", address)

		client := proto.NewCalendarClient(connection)

		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

		switch command {
		case "create":
			date, err := time.Parse("2006-01-02 15:04:05", date)
			if err != nil {
				log.Fatalf("cant parse date: %v", err)
			}

			timestamp, err := ptypes.TimestampProto(date)
			if err != nil {
				log.Fatal(err.Error())
			}

			event := &proto.Event{
				Date:        timestamp,
				Duration:    ptypes.DurationProto(time.Duration(duration) * time.Second),
				Description: description,
			}
			result, err := client.Create(ctx, event)
			if err != nil {
				log.Fatalf("cant create event: %v", err)
			}

			log.Printf("create event with id %d", result.GetId())

		case "remove":
			result, err := client.Remove(ctx, &proto.ID{Id: int64(id)})
			if err != nil {
				log.Fatalf("cant create event: %v", err)
			}
			if result.GetError() != "" {
				log.Fatalf("cant create event: %v", result.GetError())
			}
			if result.GetOk() {
				log.Printf("remove event with id %d", id)
				return
			}
			log.Fatal("cant create event by unknow reason")

		case "getlist":
			time, err := time.Parse("2006-01-02 15:04:05", date)
			if err != nil {
				log.Fatalf("cant parse date: %v", err)
			}

			timestamp, err := ptypes.TimestampProto(time)
			if err != nil {
				log.Fatal(err.Error())
			}

			result, err := client.GetActive(ctx, timestamp)
			if err != nil {
				log.Fatalf("cant get active events: %v", err)
			}

			var builder strings.Builder
			builder.WriteString(fmt.Sprintf("active events by date %s:\n", date))
			for _, event := range result.GetEvents() {
				date, err := ptypes.Timestamp(event.Date)
				if err != nil {
					log.Fatalf("cant decode timestamp: %v", err)
				}
				builder.WriteString(date.Format("2006-01-02 15:04:05"))
				builder.WriteRune(' ')
				builder.WriteString(event.Duration.String())
				builder.WriteRune(' ')
				builder.WriteString(event.Description)
				builder.WriteRune('\n')
			}

			log.Printf("active events by date %s: %+s", date, builder.String())

		default:
			log.Fatal("unknow command")
		}
	},
}
