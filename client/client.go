// uniary operation in grpc

package main

import (
	"context"
	proto "grpc/protoc"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	proto.UnimplementedExampleServer
}

var client proto.ExampleClient

func main() {

	// 	listener, tcpErr := net.Listen("tcp", ":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client = proto.NewExampleClient(conn)

	// implementing rest api
	r := gin.Default()
	r.GET("/sent-message-to-server/:message", clientConnectionServer)
	r.Run(":8000")

	// req := &proto.HelloRequest{SomeString: "hello from client"}

	// client.ServerReply(context.TODO(), req)

}

func clientConnectionServer(c *gin.Context) {
	variableName := c.Param("message")

	req := &proto.HelloRequest{SomeString: "hello from client"}

	client.ServerReply(context.TODO(), req)
	c.JSON(http.StatusOK, gin.H{
		"message": "message sent succesfully to the server " + variableName,
	})

}
