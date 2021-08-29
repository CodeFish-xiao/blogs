package main

import (
	"context"
	pb "github.com/CodeFish-xiao/blogs/gRPCAction/code/grpc-3/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
)

const (
	// ServerAddress 连接地址
	ServerAddress string = ":8546"
)

func main() {
	ServerSide()
}

func ServerSide() {
	// 连接服务器
	conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient := pb.NewServerSideClient(conn)
	// 创建发送结构体
	req := pb.ServerSideRequest{
		Name: "我来打开你啦",
	}
	//获取流
	stream, err := grpcClient.ServerSideHello(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call SayHello err: %v", err)
	}
	for n := 0; n < 5; n++ {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Conversations get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.Message)
	}
}
func ClientSide() {
	// 连接服务器
	conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient := pb.NewClientSideClient(conn)
	// 创建发送结构体

	res, err := grpcClient.ClientSideHello(context.Background())
	if err != nil {
		log.Fatalf("Call SayHello err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}
func Bidirectional() {
	// 连接服务器
	conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient := pb.NewBidirectionalClient(conn)

	stream, err := grpcClient.BidirectionalHello(context.Background())
	if err != nil {
		log.Fatalf("get BidirectionalHello stream err: %v", err)
	}

	for n := 0; n < 5; n++ {
		err := stream.Send(&pb.BidirectionalRequest{Name: "stream client rpc " + strconv.Itoa(n)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Conversations get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.Message)
	}
}
