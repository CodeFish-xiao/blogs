package main

import (
	"context"
	"github.com/CodeFish-xiao/blogs/gRPCAction/code/grpc-4/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"time"
)

const (
	// ServerAddress 连接地址
	ServerAddress string = ":8000"
)

func main() {
	var err error
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	// 指定自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))
	// 指定客户端interceptor
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	// 连接服务器
	conn, err := grpc.Dial(ServerAddress, opts...)
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

type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"token": "ok",
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {

	return false
}

// interceptor 客户端拦截器
func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	grpclog.Infof("method=%s req=%v rep=%v duration=%s error=%v\n", method, req, reply, time.Since(start), err)
	return err
}
