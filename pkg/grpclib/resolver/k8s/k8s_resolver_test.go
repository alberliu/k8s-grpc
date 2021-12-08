package k8s

import (
	"context"
	"fmt"
	"testing"

	"github.com/alberliu/k8s-grpc/pkg/pb"
	"google.golang.org/grpc"
)

func Test_resolveEndpoint(t *testing.T) {
	fmt.Println(resolveEndpoint("namespace.server_name:port"))
}

func TestClient(t *testing.T) {
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

func Test_isEqual(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{s1: []string{"1", "2"}, s2: []string{"2", "1"}},
			want: true,
		},
		{
			name: "",
			args: args{s1: []string{"1", "2"}, s2: []string{"1", "2"}},
			want: true,
		},
		{
			name: "",
			args: args{s1: []string{"1", "2"}, s2: []string{"1"}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEqual(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
