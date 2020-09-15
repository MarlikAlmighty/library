package main

import (
	"context"
	"github.com/MarlikAlmighty/library/v2/api"
	"google.golang.org/grpc"
	"log"
)

func main() {

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
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