package app

import (
	apiBooks "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/books"
	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) AddBookHandler(params apiBooks.AddBookParams) middleware.Responder {
	return middleware.NotImplemented("operation books AddBook has not yet been implemented")
}
