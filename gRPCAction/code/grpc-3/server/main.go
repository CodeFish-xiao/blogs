package main

import (
	"github.com/CodeFish-xiao/blogs/gRPCAction/code/grpc-3/pb"
	"github.com/CodeFish-xiao/blogs/gRPCAction/code/grpc-3/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	// Address 监听地址
	Address string = ":8546"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Panic("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	pb.RegisterClientSideServer(grpcServer, &services.BidirectionalService{})
	pb.RegisterServerSideServer(grpcServer, &services.ServerSideService{})
	pb.RegisterBidirectionalServer(grpcServer, &services.ClientSideService{})
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Panic("grpcServer.Serve err: %v", err)
	}
}
