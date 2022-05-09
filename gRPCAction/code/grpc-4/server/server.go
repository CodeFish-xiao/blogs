package main

import (
	"context"
	"errors"
	"github.com/CodeFish-xiao/blogs/gRPCAction/code/grpc-4/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

// HelloService 定义我们的服务
type HelloService struct {
}

// SayHello 实现SayHello方法
func (s *HelloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResp, error) {
	log.Println(req.Name)
	return &pb.HelloResp{Message: "hello ,I'm codefish "}, nil
}

const (
	// Address 监听地址
	Address string = ":8000"
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

	//普通方法：一元拦截器（grpc.UnaryInterceptor）
	var checkInterceptor grpc.UnaryServerInterceptor
	checkInterceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//拦截普通方法请求，验证Token
		err = Check(ctx)
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(checkInterceptor))
	// 在gRPC服务器注册我们的服务
	pb.RegisterHelloServer(grpcServer, &HelloService{})
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Panic("grpcServer.Serve err: %v", err)
	}
}

// Check 验证token
func Check(ctx context.Context) error {
	//从上下文中获取元数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "获取Token失败")
	}
	var (
		token string
	)
	if value, ok := md["token"]; ok {
		token = value[0]
	}
	if token != "test" {
		return errors.New("token err")
	}

	return nil
}
