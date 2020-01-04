package server

import (
	"context"
	"fmt"
	"net"

	pb "github.com/j0gger/grpc/pkg/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	certFile = "certs/cert.pem"
	keyFile  = "certs/key.pem"
	address  = "localhost:50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// Implement the server reply
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Greeting: fmt.Sprintf("Hello %s", in.GetName())}, nil
}

// RunServer is the main entry point
func RunServer() error {
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		return err
	}

	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(s, &server{})

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	return s.Serve(lis)
}
