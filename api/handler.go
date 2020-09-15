package api

import (
	"golang.org/x/net/context"
	"log"
)

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) Do(ctx context.Context, in *Request) (*Response, error) {

	log.Printf("Receive message %s", in.Str)

	return &Response{
		Library: &Library{}}, nil

}

/*
Shelf: []*Shelf{
				Id:         1,
				Number:     1,
				NameShelfs: 1,
				Book:       []*Book{
					Page: []*Page{
						Id:      1,
						Number:  1,
						Text:    "123",
						BooksId: 1,
					},
				},
			},
 */