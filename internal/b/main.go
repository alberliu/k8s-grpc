package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/alberliu/k8s-grpc/pkg/grpclib/resolver/k8s"
	"github.com/alberliu/k8s-grpc/pkg/logger"
	"github.com/alberliu/k8s-grpc/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

var aClient pb.AClient

type BServer struct{}

func (s *BServer) SayHello(ctx context.Context, in *pb.HelloRequestB) (*pb.HelloReplyB, error) {
	reply, err := aClient.SayHello(ctx, &pb.HelloRequestA{Name: "hello"})
	if err != nil {
		return nil, err
	}
	return &pb.HelloReplyB{Message: fmt.Sprintf("i im server b v3(%s) from %s", os.Getenv("POD_IP"), reply.Message)}, nil
}

func main() {
	// 初始化日志
	logger.InitLog("/log/b/")

	// 创建A服务的rpc客户端
	conn, err := grpc.DialContext(context.TODO(), k8s.GetK8STarget("default", "a", "80"), grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
	if err != nil {
		panic(err)
	}
	aClient = pb.NewAClient(conn)

	// 启动rpc服务
	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBServer(s, &BServer{})
	log.Println("rpc服务已经开启")
	s.Serve(lis)
}
