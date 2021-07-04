package app

import (
	"go.uber.org/zap"
	"html/template"
	"net/http"
	"path/filepath"

	apiHome "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/home"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

func (s *Service) HomeHandler(params apiHome.HomeParams) middleware.Responder {
	return middleware.ResponderFunc(func(rw http.ResponseWriter, _ runtime.Producer) {
		t := template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))
		if err := t.Execute(rw, nil); err != nil {
			s.Logger.Error("%s \n", zap.Error(err))
		}
		s.Logger.Info("Access",
			zap.String("ip", params.HTTPRequest.RemoteAddr),
			zap.String("ua", params.HTTPRequest.UserAgent()),
			zap.String("url", params.HTTPRequest.RequestURI))
	})
}
