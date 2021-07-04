package app

import (
	apiBookcases "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/bookcases"
	"github.com/go-openapi/runtime/middleware"
)

func (s *Service) PutBookcaseHandler(params apiBookcases.PutBookcaseParams) middleware.Responder {
	return middleware.NotImplemented("operation bookcases PutBookcase has not yet been implemented")
}
