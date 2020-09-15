package main

import (
	"fmt"
	"github.com/MarlikAlmighty/library/v2/api"
	"log"
	"net"

	"google.golang.org/grpc"
)


func main() {


	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := api.Server{}

	grpcServer := grpc.NewServer()

	api.RegisterMessageServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
