package grpclib

import (
	"context"
	"github.com/alberliu/k8s-grpc/pkg/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

type BServer1 struct{}

func (s *BServer1) SayHello(ctx context.Context, in *pb.HelloRequestB) (*pb.HelloReplyB, error) {
	return &pb.HelloReplyB{Message: "message1"}, nil
}

type BServer2 struct{}

func (s *BServer2) SayHello(ctx context.Context, in *pb.HelloRequestB) (*pb.HelloReplyB, error) {
	return &pb.HelloReplyB{Message: "message2"}, nil
}

func TestServer(t *testing.T) {
	go func() {
		lis, err := net.Listen("tcp", ":80")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		pb.RegisterBServer(s, &BServer1{})
		log.Println("rpc服务已经开启")
		s.Serve(lis)
	}()

	go func() {
		lis, err := net.Listen("tcp", ":81")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		pb.RegisterBServer(s, &BServer2{})
		log.Println("rpc服务已经开启")
		s.Serve(lis)
	}()
	select {}
}
