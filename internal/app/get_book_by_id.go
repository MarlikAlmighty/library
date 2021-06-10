package app

import (
	apiBooks "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/books"
	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) GetBookByIDHandler(params apiBooks.GetBookByIDParams) middleware.Responder {
	return middleware.NotImplemented("operation books GetBookByID has not yet been implemented")
}
