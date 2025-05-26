package main

import (
	"context"
	"fmt"

	pb "example/hello_server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("连接错误")
	}
	defer conn.Close()

	// 建立连接
	client := pb.NewSayHelloClient(conn)
	responese, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "Zhouzw"})

	fmt.Println(responese.GetResponse())
}
