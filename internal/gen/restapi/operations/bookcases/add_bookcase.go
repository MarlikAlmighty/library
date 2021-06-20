// Code generated by go-swagger; DO NOT EDIT.

package bookcases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddBookcaseHandlerFunc turns a function with the right signature into a add bookcase handler
type AddBookcaseHandlerFunc func(AddBookcaseParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddBookcaseHandlerFunc) Handle(params AddBookcaseParams) middleware.Responder {
	return fn(params)
}

// AddBookcaseHandler interface for that can handle valid add bookcase params
type AddBookcaseHandler interface {
	Handle(AddBookcaseParams) middleware.Responder
}

// NewAddBookcase creates a new http.Handler for the add bookcase operation
func NewAddBookcase(ctx *middleware.Context, handler AddBookcaseHandler) *AddBookcase {
	return &AddBookcase{Context: ctx, Handler: handler}
}

/* AddBookcase swagger:route POST /bookcases bookcases addBookcase

Add bookcase

*/
type AddBookcase struct {
	Context *middleware.Context
	Handler AddBookcaseHandler
}

func (o *AddBookcase) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddBookcaseParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}