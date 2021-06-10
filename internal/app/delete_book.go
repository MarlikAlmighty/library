package app

import (
	apiBooks "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/books"
	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) DeleteBookHandler(params apiBooks.DeleteBookParams) middleware.Responder {
	return middleware.NotImplemented("operation books DeleteBook has not yet been implemented")
}
