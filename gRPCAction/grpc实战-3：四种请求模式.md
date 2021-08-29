# gRPC应用实战：（三）gRPC四种请求模式

## 3.1 前言

gRPC主要有4种请求和响应模式，分别是简单模式(Simple RPC)、服务端流式（Server-side streaming RPC）、客户端流式（Client-side streaming
RPC）、和双向流式（Bidirectional streaming RPC）。其实好多顾名思义就可以知道相关信息：

- 简单模式：又称为一元 RPC，在上一节的时候，我们的例子就是简单模式，类似于常规的http请求，客户端发送请求，服务端响应请求
- 服务端流式：客户端发送请求到服务器，拿到一个流去读取返回的消息序列。 客户端读取返回的流，直到里面没有任何消息。
- 客户端流式：与服务端数据流模式相反，这次是客户端源源不断的向服务端发送数据流，而在发送结束后，由服务端返回一个响应。
- 双向流式：双方使用读写流去发送一个消息序列，两个流独立操作，双方可以同时发送和同时接收。

不同的调用方式往往代表着不同的应用场景，接下来我们就把剩下的三种来实操一遍：

温馨提示：以下的所有代码，都在 [这里](https://github.com/CodeFish-xiao/blogs/tree/main/gRPCAction/code/grpc-3) ,所有的pb文件都在pb包中。

## 3.2 服务端流式 RPC（Server-side streaming RPC）

服务器端流式 RPC，也就是是单向流，并代指 Server 为 Stream，Client 为普通的一元 RPC 请求。

### 3.2.1 proto

其实关键就是在服务端返回的数据前加上 `stream` 关键字

~~~protobuf
//一个为ServerSide的服务
service ServerSide {
  //一个ServerSideHello的方法
  rpc ServerSideHello (ServerSideRequest) returns (stream ServerSideResp) {}
}
~~~

然后运行 `protoc --go_out=plugins=grpc:. *.proto` 生成对应的代码。

### 3.2.2 实现服务端代码

#### 3.2.2.1 定义我们的服务

首先定义我们的服务 `ServerSideService` 并且实现`ServerSideHello`方法。
~~~go
type ServerSideService struct {
}

func (s *ServerSideService) ServerSideHello(request *pb.ServerSideRequest, server pb.ServerSide_ServerSideHelloServer) error {
	log.Println(request.Name)
	for n := 0; n < 5; n++ {
		// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes 
		err := server.Send(&pb.ServerSideResp{Message: "你好"})
		if err != nil {
			return err
		}
	}
	return nil
}
~~~

然后在 [server包](https://github.com/CodeFish-xiao/blogs/blob/main/gRPCAction/code/grpc-3/server/main.go) 中注册service

~~~go
	pb.RegisterServerSideServer(grpcServer, &services.ServerSideService{})
~~~~
