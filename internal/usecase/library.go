package usecase

import (
	"database/sql"
	"net/url"
	"os"
	"os/signal"
	"path"
	"strconv"
	"syscall"
	"time"

	"github.com/MarlikAlmighty/library/internal/repository/postgresql"

	"github.com/MarlikAlmighty/library/restapi/operations/book_id"
	"github.com/MarlikAlmighty/library/restapi/operations/shelves"
	"github.com/MarlikAlmighty/library/restapi/operations/shelves_id"

	"github.com/MarlikAlmighty/library/models"
	"github.com/MarlikAlmighty/library/restapi/operations"
	"github.com/MarlikAlmighty/library/restapi/operations/book_name"
	"github.com/MarlikAlmighty/library/restapi/operations/page_id"
	"github.com/go-openapi/runtime/middleware"
)

func ConfigureAPI(api *operations.LibraryAPI) {

	// Connection to database
	d, err := postgresql.InitialDataBase()
	if err != nil {
		api.Logger("error connect to database")
		os.Exit(0)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(3 * time.Second)
	ping := make(chan bool)

	// Check interrupt signal, check database on live
	go func() {

		defer func() {
			ticker.Stop()
			close(interrupt)
			close(ping)
			os.Exit(0)
		}()

		for {
			select {
			case <-ticker.C:
				if err := d.Conn.Ping(); err != nil {
					api.Logger("The database does not respond, closed her, and exit.")
					if err := d.Close(); err != nil {
						panic(err)
					}
					return
				}
			case <-interrupt:
				api.Logger("Received an interrupt signal, close database, and exit.")
				if err := d.Close(); err != nil {
					panic(err)
				}
				return
			}
		}
	}()

	// Get all shelves
	api.ShelvesShelvesHandler = shelves.ShelvesHandlerFunc(
		func(
			params shelves.ShelvesParams,
		) middleware.Responder {
			l, err := d.GetShelves(params.HTTPRequest.Context())
			if err != nil {
				return shelves.NewShelvesBadRequest().WithPayload(&models.Fail{
					Code:    int32(shelves.ShelvesBadRequestCode),
					Message: err.Error(),
				})
			}
			return shelves.NewShelvesOK().WithPayload(*l)
		},
	)

	// Get books from shelf by id
	api.ShelvesIDShelvesIDHandler = shelves_id.ShelvesIDHandlerFunc(
		func(
			params shelves_id.ShelvesIDParams,
		) middleware.Responder {

			u, err := url.Parse(params.HTTPRequest.RequestURI)
			if err != nil {
				return shelves_id.NewShelvesIDBadRequest().WithPayload(&models.Fail{
					Code:    int32(shelves_id.ShelvesIDBadRequestCode),
					Message: err.Error(),
				})
			}

			id, err := strconv.ParseInt(path.Base(u.Path), 10, 64)
			if err != nil {
				return shelves_id.NewShelvesIDBadRequest().WithPayload(&models.Fail{
					Code:    int32(shelves_id.ShelvesIDBadRequestCode),
					Message: "Bad request",
				})
			}

			sh, err := d.GetBooksByShelfs(params.HTTPRequest.Context(), id)
			if err != nil {
				return shelves_id.NewShelvesIDBadRequest().WithPayload(&models.Fail{
					Code:    int32(shelves_id.ShelvesIDBadRequestCode),
					Message: err.Error(),
				})
			}
			return shelves_id.NewShelvesIDOK().WithPayload(sh)
		},
	)

	// Get book by id
	api.BookIDBookIDHandler = book_id.BookIDHandlerFunc(
		func(
			params book_id.BookIDParams,
		) middleware.Responder {

			u, err := url.Parse(params.HTTPRequest.RequestURI)
			if err != nil {
				return book_id.NewBookIDBadRequest().WithPayload(&models.Fail{
					Code:    int32(book_id.BookIDBadRequestCode),
					Message: "Bad request",
				})
			}

			id, err := strconv.ParseInt(path.Base(u.Path), 10, 64)
			if err != nil {
				return book_id.NewBookIDBadRequest().WithPayload(&models.Fail{
					Code:    int32(book_id.BookIDBadRequestCode),
					Message: "Bad request",
				})
			}

			book, err := d.GetBookById(params.HTTPRequest.Context(), id)
			if err != nil {
				return book_id.NewBookIDBadRequest().WithPayload(&models.Fail{
					Code:    int32(book_id.BookIDBadRequestCode),
					Message: err.Error(),
				})
			}
			return book_id.NewBookIDOK().WithPayload(book)
		},
	)

	// Get book by name
	api.BookNameBookNameHandler = book_name.BookNameHandlerFunc(
		func(
			params book_name.BookNameParams,
		) middleware.Responder {

			u, err := url.Parse(params.HTTPRequest.RequestURI)
			if err != nil {
				return book_name.NewBookNameBadRequest().WithPayload(&models.Fail{
					Code:    int32(book_name.BookNameBadRequestCode),
					Message: "Bad request",
				})
			}

			res, err := d.GetBookByName(params.HTTPRequest.Context(), path.Base(u.Path))
			if err == sql.ErrNoRows {
				return book_name.NewBookNameBadRequest().WithPayload(&models.Fail{
					Code:    int32(book_name.BookNameBadRequestCode),
					Message: sql.ErrNoRows.Error(),
				})
			} else if err != nil {
				return book_name.NewBookNameBadRequest().WithPayload(&models.Fail{
					Code:    int32(book_name.BookNameBadRequestCode),
					Message: err.Error(),
				})
			}

			return book_name.NewBookNameOK().WithPayload(res)
		},
	)

	// Get pages by id
	api.PageIDPageIDHandler = page_id.PageIDHandlerFunc(
		func(
			params page_id.PageIDParams,
		) middleware.Responder {

			u, err := url.Parse(params.HTTPRequest.RequestURI)
			if err != nil {
				return page_id.NewPageIDBadRequest().WithPayload(&models.Fail{
					Code:    int32(page_id.PageIDBadRequestCode),
					Message: "Bad request",
				})
			}

			id, err := strconv.ParseInt(u.Path, 10, 64)
			if err != nil {
				return page_id.NewPageIDBadRequest().WithPayload(&models.Fail{
					Code:    int32(page_id.PageIDBadRequestCode),
					Message: "Bad request",
				})
			}

			pg, err := d.GetPageByID(params.HTTPRequest.Context(), id)
			if err == sql.ErrNoRows {
				return page_id.NewPageIDBadRequest().WithPayload(&models.Fail{
					Code:    int32(page_id.PageIDBadRequestCode),
					Message: sql.ErrNoRows.Error(),
				})
			} else if err != nil {
				return page_id.NewPageIDBadRequest().WithPayload(&models.Fail{
					Code:    int32(page_id.PageIDBadRequestCode),
					Message: err.Error(),
				})
			}

			return page_id.NewPageIDOK().WithPayload(pg)
		},
	)
}
