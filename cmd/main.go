package main

import (
	"log"

	"github.com/MarlikAlmighty/library/internal/gen/restapi"
	"github.com/MarlikAlmighty/library/internal/gen/restapi/operations"

	apiHome "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/home"

	apiBooks "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/books"

	apiBookcases "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/bookcases"

	"github.com/go-openapi/loads"

	"github.com/MarlikAlmighty/library/internal/app"
	"github.com/MarlikAlmighty/library/internal/config"
)

func main() {

	prefix := "LIBRARY"

	cfg, err := config.LoadEnv(prefix)
	if err != nil {
		log.Fatalln(err)
	}

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	srv := app.New()
	api := operations.NewLibraryAPI(swaggerSpec)

	api.BooksAddBookHandler = apiBooks.AddBookHandlerFunc(srv.AddBookHandler)
	api.BookcasesAddBookcaseHandler = apiBookcases.AddBookcaseHandlerFunc(srv.AddBookcaseHandler)
	api.BooksDeleteBookHandler = apiBooks.DeleteBookHandlerFunc(srv.DeleteBookHandler)
	api.BookcasesDeleteBookcaseHandler = apiBookcases.DeleteBookcaseHandlerFunc(srv.DeleteBookcaseHandler)
	api.BooksGetBookByIDHandler = apiBooks.GetBookByIDHandlerFunc(srv.GetBookByIDHandler)
	api.HomeHomeHandler = apiHome.HomeHandlerFunc(srv.HomeHandler)
	api.BookcasesListAllBookcasesHandler = apiBookcases.ListAllBookcasesHandlerFunc(srv.ListAllBookcasesHandler)
	api.BooksListAllBooksHandler = apiBooks.ListAllBooksHandlerFunc(srv.ListAllBooksHandler)
	api.BookcasesListAllBooksFromBookcasesHandler = apiBookcases.ListAllBooksFromBookcasesHandlerFunc(srv.ListAllBooksFromBookcasesHandler)
	api.BooksPutBookHandler = apiBooks.PutBookHandlerFunc(srv.PutBookHandler)
	api.BookcasesPutBookcaseHandler = apiBookcases.PutBookcaseHandlerFunc(srv.PutBookcaseHandler)

	api.ServerShutdown = srv.OnShutdown
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.ConfigureAPI()

	server.Port = cfg.HTTPPort

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
