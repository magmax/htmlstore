package main;

import (
	"fmt"
	"flag"
	"net"
	"log"
	"context"
	"google.golang.org/grpc"

	pb "htmlstore/api"
)

var (
	port = flag.Int("port", 3987, "The server port")
)

type server struct {
	pb.UnimplementedHtmlStoreServer
}

func (s *server) PutHtml(ctx context.Context, html *pb.Html) (*pb.Id, error) {
	return &pb.Id{Id: "hello!"}, nil
}


func main(){
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHtmlStoreServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
