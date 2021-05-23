// Code generated by go-docs; DO NOT EDIT.

package book_id

// This file was generated by the docs tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// BookIDHandlerFunc turns a function with the right signature into a book ID handler
type BookIDHandlerFunc func(BookIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BookIDHandlerFunc) Handle(params BookIDParams) middleware.Responder {
	return fn(params)
}

// BookIDHandler interface for that can handle valid book ID params
type BookIDHandler interface {
	Handle(BookIDParams) middleware.Responder
}

// NewBookID creates a new http.Handler for the book ID operation
func NewBookID(ctx *middleware.Context, handler BookIDHandler) *BookID {
	return &BookID{Context: ctx, Handler: handler}
}

/*BookID docs:route GET /book/id/{id} bookID bookId

Handler for seek book on id

*/
type BookID struct {
	Context *middleware.Context
	Handler BookIDHandler
}

func (o *BookID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBookIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
