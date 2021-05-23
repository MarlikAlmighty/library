// Code generated by go-docs; DO NOT EDIT.

package shelves

// This file was generated by the docs tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ShelvesHandlerFunc turns a function with the right signature into a shelves handler
type ShelvesHandlerFunc func(ShelvesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ShelvesHandlerFunc) Handle(params ShelvesParams) middleware.Responder {
	return fn(params)
}

// ShelvesHandler interface for that can handle valid shelves params
type ShelvesHandler interface {
	Handle(ShelvesParams) middleware.Responder
}

// NewShelves creates a new http.Handler for the shelves operation
func NewShelves(ctx *middleware.Context, handler ShelvesHandler) *Shelves {
	return &Shelves{Context: ctx, Handler: handler}
}

/*Shelves docs:route GET /shelves shelves shelves

Handler for show list shelves

*/
type Shelves struct {
	Context *middleware.Context
	Handler ShelvesHandler
}

func (o *Shelves) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewShelvesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
