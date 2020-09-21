package main

import (
	"context"
	"github.com/MarlikAlmighty/library/v2/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {

	var conn *grpc.ClientConn

	creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	conn, err = grpc.Dial("localhost:7777", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewMessageClient(conn)

	response, err := c.Do(context.Background(), &api.Request{Str: "foo"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}


	log.Printf("Response from server: %s", response)
}