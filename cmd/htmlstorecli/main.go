package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "htmlstore/api"
)

var (
	addr = flag.String("addr", "localhost:3987", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewHtmlStoreClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if id, err := c.PutHtml(ctx, &pb.Html{Content: "hello"}); err != nil {
		log.Fatalf("could not greet: %v", err)
	} else {
		log.Fatalf("Inserted as %s", id)
	}
}
