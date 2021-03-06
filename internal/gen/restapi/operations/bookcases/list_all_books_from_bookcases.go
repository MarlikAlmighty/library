// Code generated by go-swagger; DO NOT EDIT.

package bookcases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAllBooksFromBookcasesHandlerFunc turns a function with the right signature into a list all books from bookcases handler
type ListAllBooksFromBookcasesHandlerFunc func(ListAllBooksFromBookcasesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAllBooksFromBookcasesHandlerFunc) Handle(params ListAllBooksFromBookcasesParams) middleware.Responder {
	return fn(params)
}

// ListAllBooksFromBookcasesHandler interface for that can handle valid list all books from bookcases params
type ListAllBooksFromBookcasesHandler interface {
	Handle(ListAllBooksFromBookcasesParams) middleware.Responder
}

// NewListAllBooksFromBookcases creates a new http.Handler for the list all books from bookcases operation
func NewListAllBooksFromBookcases(ctx *middleware.Context, handler ListAllBooksFromBookcasesHandler) *ListAllBooksFromBookcases {
	return &ListAllBooksFromBookcases{Context: ctx, Handler: handler}
}

/* ListAllBooksFromBookcases swagger:route GET /bookcases/{id} bookcases listAllBooksFromBookcases

List all books in bookcase

*/
type ListAllBooksFromBookcases struct {
	Context *middleware.Context
	Handler ListAllBooksFromBookcasesHandler
}

func (o *ListAllBooksFromBookcases) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListAllBooksFromBookcasesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
