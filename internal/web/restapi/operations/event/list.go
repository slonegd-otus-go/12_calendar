// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListHandlerFunc turns a function with the right signature into a list handler
type ListHandlerFunc func(ListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListHandlerFunc) Handle(params ListParams) middleware.Responder {
	return fn(params)
}

// ListHandler interface for that can handle valid list params
type ListHandler interface {
	Handle(ListParams) middleware.Responder
}

// NewList creates a new http.Handler for the list operation
func NewList(ctx *middleware.Context, handler ListHandler) *List {
	return &List{Context: ctx, Handler: handler}
}

/*List swagger:route GET /events event list

List list API

*/
type List struct {
	Context *middleware.Context
	Handler ListHandler
}

func (o *List) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
