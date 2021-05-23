package postgresql

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/MarlikAlmighty/library/models"

	// nolint: unconvert
	_ "github.com/lib/pq"
)

// Handler for pool connections
type Handler struct {
	Conn *sql.DB
}

// InitialDataBase database connection
func InitialDataBase() (*Handler, error) {

	db, err := sql.Open("postgres", os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		return &Handler{Conn: nil}, err
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)

	if err := db.Ping(); err != nil {
		return &Handler{Conn: nil}, err
	}

	return &Handler{Conn: db}, nil
}

func (p *Handler) Close() error {
	if err := p.Conn.Close(); err != nil {
		return err
	}
	return nil
}

// GetPageByID return page by id
func (p *Handler) GetPageByID(ctx context.Context, id int64) (*models.Page, error) {

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	m := models.Page{}

	row := p.Conn.QueryRowContext(ctx, `SELECT pages_id, books_id, pages_number, pages_text 
	FROM pages WHERE pages_id = $1;`, id)

	err := row.Scan(&m.ID, &m.BooksID, &m.Number, &m.Text)
	if err != nil {
		return &m, err
	}

	return &m, err
}

// GetBookByName return book by name
func (p *Handler) GetBookByName(ctx context.Context, name string) (*models.Book, error) {

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	b := models.Book{}

	row := p.Conn.QueryRowContext(ctx, `SELECT books_id, author, num_pages, publisher, shelfs_id, pages_id 
	FROM books WHERE author = $1;`, name)

	err := row.Scan(&b.ID, &b.Author, &b.NumPages, &b.Publisher, &b.ShelvesID, &b.Pages)
	if err != nil {
		return &b, err
	}

	return &b, err
}

// GetBookById return book by id
func (p *Handler) GetBookById(ctx context.Context, id int64) (*models.Book, error) {

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	b := models.Book{}

	row := p.Conn.QueryRowContext(ctx, `SELECT books_id, author, num_pages, publisher, shelfs_id
	FROM books WHERE books_id = $1;`, id)

	err := row.Scan(&b.ID, &b.Author, &b.NumPages, &b.Publisher, &b.ShelvesID)
	if err != nil {
		return &b, err
	}

	return &b, err
}

// GetShelfsID return books by id of shelf
func (p *Handler) GetBooksByShelfs(ctx context.Context, id int64) (*models.Shelf, error) {

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	s := models.Shelf{}

	rows, err := p.Conn.QueryContext(ctx, `SELECT books_id, author, num_pages, publisher, shelfs_id
	FROM books WHERE shelfs_id = $1;`, id)
	if err != nil {
		return &s, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			panic(err)
		}
	}()

	for rows.Next() {

		b := models.Book{}

		err = rows.Scan(&b.ID, &b.Author, &b.NumPages, &b.Publisher, &b.ShelvesID)
		if err != nil {
			return &s, err
		}

		s.Books = append(s.Books, &b)
	}

	if err = rows.Err(); err != nil {
		return &s, err
	}

	return &s, nil
}

// GetShelves return all shelves
func (p *Handler) GetShelves(ctx context.Context) (*models.Library, error) {

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	l := models.Library{}

	rows, err := p.Conn.QueryContext(ctx, `SELECT shelfs_id, shelfs_name, shelfs_number FROM shelfs;`)
	if err != nil {
		return &l, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			panic(err)
		}
	}()

	for rows.Next() {

		s := models.Shelf{}

		err = rows.Scan(&s.ID, &s.NameShelfs, &s.Number)
		if err != nil {
			return &l, err
		}

		l = append(l, &s)

	}

	if err = rows.Err(); err != nil {
		return &l, err
	}

	return &l, nil
}