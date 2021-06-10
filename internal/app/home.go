package app

import (
	apiHome "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/home"
	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) HomeHandler(params apiHome.HomeParams) middleware.Responder {
	return middleware.NotImplemented("operation home Home has not yet been implemented")
}
