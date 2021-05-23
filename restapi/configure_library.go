// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	usecase2 "github.com/MarlikAlmighty/library/internal/usecase"
	"log"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/MarlikAlmighty/library/restapi/operations"
	"github.com/MarlikAlmighty/library/restapi/operations/book_id"
	"github.com/MarlikAlmighty/library/restapi/operations/book_name"
	"github.com/MarlikAlmighty/library/restapi/operations/page_id"
	"github.com/MarlikAlmighty/library/restapi/operations/shelves"
	"github.com/MarlikAlmighty/library/restapi/operations/shelves_id"
)

//go:generate swagger generate server --target ../../usecase --name Library --spec ../swagger/swagger.yml

func configureFlags(api *operations.LibraryAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.LibraryAPI) http.Handler {

	api.ServeError = errors.ServeError

	api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.BookIDBookIDHandler == nil {
		api.BookIDBookIDHandler = book_id.BookIDHandlerFunc(func(params book_id.BookIDParams) middleware.Responder {
			return middleware.NotImplemented("operation book_id.BookID has not yet been implemented")
		})
	}
	if api.BookNameBookNameHandler == nil {
		api.BookNameBookNameHandler = book_name.BookNameHandlerFunc(func(params book_name.BookNameParams) middleware.Responder {
			return middleware.NotImplemented("operation book_name.BookName has not yet been implemented")
		})
	}
	if api.PageIDPageIDHandler == nil {
		api.PageIDPageIDHandler = page_id.PageIDHandlerFunc(func(params page_id.PageIDParams) middleware.Responder {
			return middleware.NotImplemented("operation page_id.PageID has not yet been implemented")
		})
	}
	if api.ShelvesShelvesHandler == nil {
		api.ShelvesShelvesHandler = shelves.ShelvesHandlerFunc(func(params shelves.ShelvesParams) middleware.Responder {
			return middleware.NotImplemented("operation shelves.Shelves has not yet been implemented")
		})
	}
	if api.ShelvesIDShelvesIDHandler == nil {
		api.ShelvesIDShelvesIDHandler = shelves_id.ShelvesIDHandlerFunc(func(params shelves_id.ShelvesIDParams) middleware.Responder {
			return middleware.NotImplemented("operation shelves_id.ShelvesID has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	usecase2.ConfigureAPI(api)

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the docs.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the docs.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
