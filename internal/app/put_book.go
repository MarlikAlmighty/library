package app

import (
	apiBooks "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/books"
	"github.com/go-openapi/runtime/middleware"
)

func (s *Service) PutBookHandler(params apiBooks.PutBookParams) middleware.Responder {
	return middleware.NotImplemented("operation books PutBook has not yet been implemented")
}
