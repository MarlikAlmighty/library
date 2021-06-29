package main

import (
	"github.com/MarlikAlmighty/library/internal/app"
	"github.com/MarlikAlmighty/library/internal/gen/restapi"
	"github.com/MarlikAlmighty/library/internal/gen/restapi/operations"
	"log"

	apiHome "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/home"

	apiBooks "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/books"

	apiBookcases "github.com/MarlikAlmighty/library/internal/gen/restapi/operations/bookcases"

	"github.com/go-openapi/loads"
)

func main() {

	myApp, err := app.New()
	if err != nil {
		log.Fatalln(err)
	}

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewLibraryAPI(swaggerSpec)

	api.BooksAddBookHandler = apiBooks.AddBookHandlerFunc(myApp.AddBookHandler)
	api.BookcasesAddBookcaseHandler = apiBookcases.AddBookcaseHandlerFunc(myApp.AddBookcaseHandler)
	api.BooksDeleteBookHandler = apiBooks.DeleteBookHandlerFunc(myApp.DeleteBookHandler)
	api.BookcasesDeleteBookcaseHandler = apiBookcases.DeleteBookcaseHandlerFunc(myApp.DeleteBookcaseHandler)
	api.BooksGetBookByIDHandler = apiBooks.GetBookByIDHandlerFunc(myApp.GetBookByIDHandler)
	api.HomeHomeHandler = apiHome.HomeHandlerFunc(myApp.HomeHandler)
	api.BookcasesListAllBookcasesHandler = apiBookcases.ListAllBookcasesHandlerFunc(myApp.ListAllBookcasesHandler)
	api.BooksListAllBooksHandler = apiBooks.ListAllBooksHandlerFunc(myApp.ListAllBooksHandler)
	api.BookcasesListAllBooksFromBookcasesHandler = apiBookcases.ListAllBooksFromBookcasesHandlerFunc(myApp.ListAllBooksFromBookcasesHandler)
	api.BooksPutBookHandler = apiBooks.PutBookHandlerFunc(myApp.PutBookHandler)
	api.BookcasesPutBookcaseHandler = apiBookcases.PutBookcaseHandlerFunc(myApp.PutBookcaseHandler)

	server := restapi.NewServer(api)

	defer func() {
		myApp.Stop()
		if err := server.Shutdown(); err != nil {
			log.Fatalln(err)
		}
	}()

	server.ConfigureAPI()

	server.Port = myApp.Conf.HTTPPort

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
