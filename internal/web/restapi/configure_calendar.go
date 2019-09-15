// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/slonegd-otus-go/12_calendar/internal/web/restapi/operations"
	"github.com/slonegd-otus-go/12_calendar/internal/web/restapi/operations/event"
)

//go:generate swagger generate server --target ../../web --name Calendar --spec ../../../api/swagger.yml --exclude-main

func configureFlags(api *operations.CalendarAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CalendarAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.EventCreateHandler == nil {
		api.EventCreateHandler = event.CreateHandlerFunc(func(params event.CreateParams) middleware.Responder {
			return middleware.NotImplemented("operation event.Create has not yet been implemented")
		})
	}
	if api.EventGetHandler == nil {
		api.EventGetHandler = event.GetHandlerFunc(func(params event.GetParams) middleware.Responder {
			return middleware.NotImplemented("operation event.Get has not yet been implemented")
		})
	}
	if api.EventListHandler == nil {
		api.EventListHandler = event.ListHandlerFunc(func(params event.ListParams) middleware.Responder {
			return middleware.NotImplemented("operation event.List has not yet been implemented")
		})
	}
	if api.EventRemoveHandler == nil {
		api.EventRemoveHandler = event.RemoveHandlerFunc(func(params event.RemoveParams) middleware.Responder {
			return middleware.NotImplemented("operation event.Remove has not yet been implemented")
		})
	}
	if api.EventUpdateHandler == nil {
		api.EventUpdateHandler = event.UpdateHandlerFunc(func(params event.UpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation event.Update has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
