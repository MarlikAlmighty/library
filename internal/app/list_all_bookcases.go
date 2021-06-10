package app

import (
	apiBookcases "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/bookcases"
	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) ListAllBookcasesHandler(params apiBookcases.ListAllBookcasesParams) middleware.Responder {
	return middleware.NotImplemented("operation bookcases ListAllBookcases has not yet been implemented")
}
