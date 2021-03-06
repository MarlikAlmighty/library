// Code generated by go-swagger; DO NOT EDIT.

package bookcases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutBookcaseHandlerFunc turns a function with the right signature into a put bookcase handler
type PutBookcaseHandlerFunc func(PutBookcaseParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutBookcaseHandlerFunc) Handle(params PutBookcaseParams) middleware.Responder {
	return fn(params)
}

// PutBookcaseHandler interface for that can handle valid put bookcase params
type PutBookcaseHandler interface {
	Handle(PutBookcaseParams) middleware.Responder
}

// NewPutBookcase creates a new http.Handler for the put bookcase operation
func NewPutBookcase(ctx *middleware.Context, handler PutBookcaseHandler) *PutBookcase {
	return &PutBookcase{Context: ctx, Handler: handler}
}

/* PutBookcase swagger:route PUT /bookcases bookcases putBookcase

Update an existing bookcase

*/
type PutBookcase struct {
	Context *middleware.Context
	Handler PutBookcaseHandler
}

func (o *PutBookcase) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutBookcaseParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
