package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/alberliu/k8s-grpc/pkg/logger"
	"github.com/alberliu/k8s-grpc/pkg/pb"
	"google.golang.org/grpc"
)

type AServer struct{}

func (s *AServer) SayHello(ctx context.Context, in *pb.HelloRequestA) (*pb.HelloReplyA, error) {
	return &pb.HelloReplyA{Message: fmt.Sprintf("i im server a v3(%s)", os.Getenv("POD_IP"))}, nil
}

func main() {
	logger.InitLog("/log/a/")

	// 启动rpc服务
	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAServer(s, &AServer{})
	log.Println("rpc服务已经开启")
	s.Serve(lis)
}
