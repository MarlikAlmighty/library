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