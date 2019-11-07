package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/slonegd-otus-go/12_calendar/internal/event"
)

type Server struct {
	storage event.Storage
}

func NewServer(storage event.Storage) *Server {
	return &Server{storage}
}

func (server *Server) Create(_ context.Context, eventRequest *Event) (*Event, error) {
	date := time.Time{}
	if eventRequest.GetDate() != nil {
		tmp, err := ptypes.Timestamp(eventRequest.GetDate())
		if err != nil {
			return nil, err
		}
		date = tmp
	}

	var duration time.Duration
	if eventRequest.GetDuration() != nil {
		tmp, err := ptypes.Duration(eventRequest.GetDuration())
		if err != nil {
			return nil, err
		}
		duration = tmp
	}

	id := server.storage.Add(event.Event{
		Date:        date,
		Duration:    duration,
		Description: eventRequest.GetDescription(),
	})
	eventRequest.Id = int64(id)

	return eventRequest, nil

}

func (server *Server) GetActive(_ context.Context, dateRequest *timestamp.Timestamp) (*Events, error) {
	date, err := ptypes.Timestamp(dateRequest)
	if err != nil {
		return nil, err
	}

	active := server.storage.Active(date)

	events := &Events{}
	for id, event := range active {
		duration := ptypes.DurationProto(event.Duration)
		date, err := ptypes.TimestampProto(event.Date)
		if err != nil {
			return nil, fmt.Errorf("parse TimestampProto from time.Time failed: %s", err)
		}
		events.Events = append(events.Events, &Event{
			Id:          int64(id),
			Date:        date,
			Duration:    duration,
			Description: event.Description,
		})
	}

	return events, nil
}

func (server *Server) Get(_ context.Context, id *ID) (*GetResponse, error) {
	event, ok := server.storage.Get(event.ID(id.Id))

	if !ok {
		return &GetResponse{Result: &GetResponse_Error{
			fmt.Sprintf("dont have event for id %d", id.Id),
		}}, nil
	}

	date, err := ptypes.TimestampProto(event.Date)
	if err != nil {
		return nil, fmt.Errorf("parse TimestampProto from time.Time failed: %s", err)
	}
	duration := ptypes.DurationProto(event.Duration)

	return &GetResponse{Result: &GetResponse_Event{&Event{
		Id:          id.Id,
		Date:        date,
		Duration:    duration,
		Description: event.Description,
	}}}, nil
}

func (server *Server) Remove(_ context.Context, id *ID) (*ChangeResponse, error) {
	ok := server.storage.Remove(event.ID(id.Id))

	if !ok {
		return &ChangeResponse{Result: &ChangeResponse_Error{
			fmt.Sprintf("dont have event for id %d", id.Id),
		}}, nil
	}

	return &ChangeResponse{Result: &ChangeResponse_Ok{ok}}, nil
}

func (server *Server) Update(_ context.Context, eventRequest *Event) (*ChangeResponse, error) {
	date := time.Time{}
	if eventRequest.GetDate() != nil {
		tmp, err := ptypes.Timestamp(eventRequest.GetDate())
		if err != nil {
			return nil, err
		}
		date = tmp
	}

	var duration time.Duration
	if eventRequest.GetDuration() != nil {
		tmp, err := ptypes.Duration(eventRequest.GetDuration())
		if err != nil {
			return nil, err
		}
		duration = tmp
	}

	id := event.ID(eventRequest.GetId())

	ok := server.storage.Update(id, event.Event{
		Date:        date,
		Duration:    duration,
		Description: eventRequest.GetDescription(),
	})

	if !ok {
		return &ChangeResponse{Result: &ChangeResponse_Error{
			fmt.Sprintf("dont have event for id %d", id),
		}}, nil
	}

	return &ChangeResponse{Result: &ChangeResponse_Ok{ok}}, nil
}
