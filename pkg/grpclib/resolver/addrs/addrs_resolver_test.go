package addrs

import (
	"context"
	"fmt"
	"testing"

	"github.com/alberliu/k8s-grpc/pkg/pb"
	"google.golang.org/grpc"
)

func TestAddrsResolver(t *testing.T) {
	conn, err := grpc.DialContext(context.TODO(), "127.0.0.1:30080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewBClient(conn)
	fmt.Println(client.SayHello(context.TODO(), &pb.HelloRequestB{Name: "hello"}))
	fmt.Println(client.SayHello(context.TODO(), &pb.HelloRequestB{Name: "hello"}))
	fmt.Println(client.SayHello(context.TODO(), &pb.HelloRequestB{Name: "hello"}))
	fmt.Println(client.SayHello(context.TODO(), &pb.HelloRequestB{Name: "hello"}))
	fmt.Println(client.SayHello(context.TODO(), &pb.HelloRequestB{Name: "hello"}))
	fmt.Println(client.SayHello(context.TODO(), &pb.HelloRequestB{Name: "hello"}))
	fmt.Println(client.SayHello(context.TODO(), &pb.HelloRequestB{Name: "hello"}))
	fmt.Println(client.SayHello(context.TODO(), &pb.HelloRequestB{Name: "hello"}))
}

func TestAddrsResolver1(t *testing.T) {
	conn, err := grpc.DialContext(context.TODO(), "127.0.0.1:30080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewAClient(conn)
	fmt.Println(client.SayHello(context.TODO(), &pb.HelloRequestA{Name: "hello"}))
}
