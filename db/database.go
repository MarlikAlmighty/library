package db

import (
	"context"
	"database/sql"
	"os"
	"time"

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

