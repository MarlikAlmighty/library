package app

import (
	apiBookcases "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/bookcases"
	"github.com/go-openapi/runtime/middleware"
)

func (s *Service) ListAllBooksFromBookcasesHandler(params apiBookcases.ListAllBooksFromBookcasesParams) middleware.Responder {
	return middleware.NotImplemented("operation bookcases ListAllBooksFromBookcases has not yet been implemented")
}
