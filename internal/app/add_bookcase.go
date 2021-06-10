package app

import (
	apiBookcases "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/bookcases"
	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) AddBookcaseHandler(params apiBookcases.AddBookcaseParams) middleware.Responder {
	return middleware.NotImplemented("operation bookcases AddBookcase has not yet been implemented")
}
