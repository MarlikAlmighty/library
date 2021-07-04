package app

import (
	apiBooks "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/books"
	"github.com/go-openapi/runtime/middleware"
)

func (s *Service) ListAllBooksHandler(params apiBooks.ListAllBooksParams) middleware.Responder {
	return middleware.NotImplemented("operation books ListAllBooks has not yet been implemented")
}
