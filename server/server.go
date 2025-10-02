// uniary operation in grpc

package main

import (
	"context"
	"errors"
	"fmt"
	proto "grpc/protoc"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExampleServer
}

func main() {

	// 	listener, tcpErr := net.Listen("tcp", ":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}

	srv := grpc.NewServer()
	proto.RegisterExampleServer(srv, &server{})
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) ServerReply(c context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println("Recieve request from client: ", req.SomeString)
	fmt.Println("hello from Server!!!")
	return &proto.HelloResponse{}, errors.New("")
}
