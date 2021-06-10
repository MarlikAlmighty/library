package app

import (
	apiBookcases "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/bookcases"
	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) DeleteBookcaseHandler(params apiBookcases.DeleteBookcaseParams) middleware.Responder {
	return middleware.NotImplemented("operation bookcases DeleteBookcase has not yet been implemented")
}
