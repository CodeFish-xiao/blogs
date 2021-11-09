package main

import (
	"context"
	"github.com/CodeFish-xiao/blogs/gRPCAction/code/grpc-4/pb"
	"google.golang.org/grpc"
	"log"
)

const (
	// ServerAddress 连接地址
	ServerAddress string = ":8000"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient := pb.NewHelloClient(conn)
	// 创建发送结构体
	req := pb.HelloRequest{
		Name: "grpc",
	}
	// 调用我们的服务(SayHello方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := grpcClient.SayHello(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call SayHello err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}
