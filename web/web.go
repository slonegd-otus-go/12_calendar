package web

import (
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/slonegd-otus-go/12_calendar/internal/event"
	"github.com/slonegd-otus-go/12_calendar/web/models"
	"github.com/slonegd-otus-go/12_calendar/web/restapi"
	"github.com/slonegd-otus-go/12_calendar/web/restapi/operations"
	eventapi "github.com/slonegd-otus-go/12_calendar/web/restapi/operations/event"
)

func Run(port int, storage *event.Storage) {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewCalendarAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = port

	api.EventCreateHandler = eventapi.CreateHandlerFunc(
		func(params eventapi.CreateParams) middleware.Responder {
			return middleware.NotImplemented("operation event.Create has not yet been implemented")
		})

	api.EventGetHandler = eventapi.GetHandlerFunc(
		func(params eventapi.GetParams) middleware.Responder {
			return middleware.NotImplemented("operation event.Get has not yet been implemented")
		})

	api.EventListHandler = eventapi.ListHandlerFunc(
		func(params eventapi.ListParams) middleware.Responder {
			var payload []*models.Event
			storage.Range(func(id event.ID, event event.Event) {
				date := event.Date.Format("2006-01-02 15:04:05")
				duration := int64(event.Duration.Seconds())
				payload = append(payload, &models.Event{
					ID:          int64(id),
					Date:        &date,
					Duration:    &duration,
					Description: &event.Description,
				})
			})
			return eventapi.NewListOK().WithPayload(payload)
		})

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
