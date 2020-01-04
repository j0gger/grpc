package client

import (
	"context"
	"os"
	"time"

	pb "github.com/j0gger/grpc/pkg/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address     = "localhost:50051"
	defaultName = "acme"
)

// MakeRequest to the server
func MakeRequest() (string, error) {
	creds, err := credentials.NewClientTLSFromFile("certs/cert.pem", "")
	if err != nil {
		return "", err
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		return "", err
	}

	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 2 {
		name = os.Args[2]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})

	if err != nil {
		return "", err
	}

	return r.GetGreeting(), nil
}
