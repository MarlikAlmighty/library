// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/MarlikAlmighty/library/internal/gen/restapi/operations"
	"github.com/MarlikAlmighty/library/internal/gen/restapi/operations/bookcases"
	"github.com/MarlikAlmighty/library/internal/gen/restapi/operations/books"
	"github.com/MarlikAlmighty/library/internal/gen/restapi/operations/home"
)

//go:generate swagger generate server --target ../../gen --name Library --spec ../../../swagger-api/swagger.yml --template-dir ./swagger-templates --principal interface{}

func configureFlags(api *operations.LibraryAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.LibraryAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.HTMLProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("html producer has not yet been implemented")
	})
	api.JSONProducer = runtime.JSONProducer()

	if api.BooksAddBookHandler == nil {
		api.BooksAddBookHandler = books.AddBookHandlerFunc(func(params books.AddBookParams) middleware.Responder {
			return middleware.NotImplemented("operation books.AddBook has not yet been implemented")
		})
	}
	if api.BookcasesAddBookcaseHandler == nil {
		api.BookcasesAddBookcaseHandler = bookcases.AddBookcaseHandlerFunc(func(params bookcases.AddBookcaseParams) middleware.Responder {
			return middleware.NotImplemented("operation bookcases.AddBookcase has not yet been implemented")
		})
	}
	if api.BooksDeleteBookHandler == nil {
		api.BooksDeleteBookHandler = books.DeleteBookHandlerFunc(func(params books.DeleteBookParams) middleware.Responder {
			return middleware.NotImplemented("operation books.DeleteBook has not yet been implemented")
		})
	}
	if api.BookcasesDeleteBookcaseHandler == nil {
		api.BookcasesDeleteBookcaseHandler = bookcases.DeleteBookcaseHandlerFunc(func(params bookcases.DeleteBookcaseParams) middleware.Responder {
			return middleware.NotImplemented("operation bookcases.DeleteBookcase has not yet been implemented")
		})
	}
	if api.BooksGetBookByIDHandler == nil {
		api.BooksGetBookByIDHandler = books.GetBookByIDHandlerFunc(func(params books.GetBookByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation books.GetBookByID has not yet been implemented")
		})
	}
	if api.HomeHomeHandler == nil {
		api.HomeHomeHandler = home.HomeHandlerFunc(func(params home.HomeParams) middleware.Responder {
			return middleware.NotImplemented("operation home.Home has not yet been implemented")
		})
	}
	if api.BookcasesListAllBookcasesHandler == nil {
		api.BookcasesListAllBookcasesHandler = bookcases.ListAllBookcasesHandlerFunc(func(params bookcases.ListAllBookcasesParams) middleware.Responder {
			return middleware.NotImplemented("operation bookcases.ListAllBookcases has not yet been implemented")
		})
	}
	if api.BooksListAllBooksHandler == nil {
		api.BooksListAllBooksHandler = books.ListAllBooksHandlerFunc(func(params books.ListAllBooksParams) middleware.Responder {
			return middleware.NotImplemented("operation books.ListAllBooks has not yet been implemented")
		})
	}
	if api.BookcasesListAllBooksFromBookcasesHandler == nil {
		api.BookcasesListAllBooksFromBookcasesHandler = bookcases.ListAllBooksFromBookcasesHandlerFunc(func(params bookcases.ListAllBooksFromBookcasesParams) middleware.Responder {
			return middleware.NotImplemented("operation bookcases.ListAllBooksFromBookcases has not yet been implemented")
		})
	}
	if api.BooksPutBookHandler == nil {
		api.BooksPutBookHandler = books.PutBookHandlerFunc(func(params books.PutBookParams) middleware.Responder {
			return middleware.NotImplemented("operation books.PutBook has not yet been implemented")
		})
	}
	if api.BookcasesPutBookcaseHandler == nil {
		api.BookcasesPutBookcaseHandler = bookcases.PutBookcaseHandlerFunc(func(params bookcases.PutBookcaseParams) middleware.Responder {
			return middleware.NotImplemented("operation bookcases.PutBookcase has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
