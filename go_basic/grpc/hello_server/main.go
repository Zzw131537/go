package main

import (
	"context"
	pb "example/hello_server/proto"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

// 自定义hello
type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Response: "hello" + req.RequestName}, nil
}

func main() {
	// 开启端口
	listen, _ := net.Listen("tcp", ":9090")
	grpcServer := grpc.NewServer()
	// 注册服务
	pb.RegisterSayHelloServer(grpcServer, &server{})

	// 启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Println("启动服务失败")
	}
}
