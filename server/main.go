package main

import (
	"context"
	"log"
	"net"

	pb "github.com/sakura1116vg/try_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server 構造体
// この構造体にprotoファイルで定期したRPCのメソッドを実装していく
type server struct{}

// SayHello メソッド
// protoファイルのservice部分に記載したRPCメソッドを実装する
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	//return pb.HelloReply{Message: "Hello " + in.Name}, nil
	return &pb.HelloReply{Message: in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// server構造体をGreeterとしてProtocol Bufferに登録する
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
