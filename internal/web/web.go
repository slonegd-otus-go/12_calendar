package web

import (
	"log"
	"time"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/slonegd-otus-go/12_calendar/internal/event"
	"github.com/slonegd-otus-go/12_calendar/internal/web/models"
	"github.com/slonegd-otus-go/12_calendar/internal/web/restapi"
	"github.com/slonegd-otus-go/12_calendar/internal/web/restapi/operations"
	eventapi "github.com/slonegd-otus-go/12_calendar/internal/web/restapi/operations/event"
)

func Run(host string, port int, storage event.Storage) {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewCalendarAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Host = host
	server.Port = port

	api.JSONProducer = JSONProducer{}
	api.JSONConsumer = JSONConsumer{}

	api.EventCreateHandler = eventapi.CreateHandlerFunc(
		func(params eventapi.CreateParams) middleware.Responder {
			date, err := time.Parse("2006-01-02 15:04:05", *params.Event.Date)
			if err != nil {
				return eventapi.NewCreateBadRequest()
			}

			duration := time.Duration(*params.Event.Duration) * time.Second
			id := storage.Add(event.Event{
				Date:        date,
				Duration:    duration,
				Description: *params.Event.Description,
			})
			params.Event.ID = int64(id)

			return eventapi.NewCreateOK().WithPayload(params.Event)
		})

	api.EventGetHandler = eventapi.GetHandlerFunc(
		func(params eventapi.GetParams) middleware.Responder {
			resultEvent, ok := storage.Get(event.ID(params.EventID))
			if !ok {
				return eventapi.NewGetNotFound()
			}
			date := resultEvent.Date.Format("2006-01-02 15:04:05")
			duration := int64(resultEvent.Duration.Seconds())
			return eventapi.NewGetOK().WithPayload(&models.Event{
				ID:          params.EventID,
				Date:        &date,
				Duration:    &duration,
				Description: &resultEvent.Description,
			})
		})

	api.EventListHandler = eventapi.ListHandlerFunc(
		func(params eventapi.ListParams) middleware.Responder {
			respDate, err := time.Parse("2006-01-02 15:04:05", params.Date)
			if err != nil {
				return eventapi.NewListBadRequest()
			}

			active := storage.Active(respDate)

			var events []*models.Event
			for id, event := range active {
				date := event.Date.Format("2006-01-02 15:04:05")
				duration := int64(event.Duration.Seconds())
				events = append(events, &models.Event{
					ID:          int64(id),
					Date:        &date,
					Duration:    &duration,
					Description: &event.Description,
				})
			}
			return eventapi.NewListOK().WithPayload(events)
		})

	api.EventRemoveHandler = eventapi.RemoveHandlerFunc(
		func(params eventapi.RemoveParams) middleware.Responder {
			ok := storage.Remove(event.ID(params.ID))
			if ok {
				return eventapi.NewRemoveOK()
			}
			return eventapi.NewRemoveNotFound()
		})

	api.EventUpdateHandler = eventapi.UpdateHandlerFunc(
		func(params eventapi.UpdateParams) middleware.Responder {
			date, err := time.Parse("2006-01-02 15:04:05", *params.Event.Date)
			if err != nil {
				return eventapi.NewUpdateBadRequest()
			}

			duration := time.Duration(*params.Event.Duration) * time.Second
			ok := storage.Update(event.ID(params.ID), event.Event{
				Date:        date,
				Duration:    duration,
				Description: *params.Event.Description,
			})
			params.Event.ID = params.ID
			if ok {
				return eventapi.NewUpdateOK()
			}
			return eventapi.NewUpdateNotFound()
		})

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
