package main

import (
	"context"
	pb "github.com/sakura1116vg/try_grpc"
	"google.golang.org/grpc"
	"log"
	"os"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// gRPCサーバとのコネクションを作成している
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	// Contact the server and print out its response.
	name := defaultName

	if len(os.Args) == 2 {
		name = os.Args[1]
	}
	// gRPCサーバのSayHelloメソッドを呼び出している
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
