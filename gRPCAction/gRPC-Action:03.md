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

#### 3.2.2.2 运行我们的服务

这是全部的server的信息，后面就不再重复这一部分的信息了，通过 `grpc.NewServer()` 创建新的gRPC服务器，之后进行对应的服务注册，并且调用 `grpc.NewServer()` 阻塞线程。

~~~go
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

~~~

运行后：

[![4pHWSf.png](https://z3.ax1x.com/2021/09/12/4pHWSf.png)](https://imgtu.com/i/4pHWSf)

### 3.2.2 实现客户端代码

代码如下：

~~~go
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
~~~

因为是服务器流模式，需要先从服务器获取流，也就是链接，通过流进行数据传输，客户端通过 `Recv()` 获取服务端的信息，然后输出

### 3.2.3 服务流模式运行样例

编写完客户端代码后，运行可见： 发送一次请求后，收取服务端发来的请求信息

客户端：
[![4pqoin.png](https://z3.ax1x.com/2021/09/12/4pqoin.png)](https://imgtu.com/i/4pqoin)

服务端：

收到一次客户端的请求后，发送信息。

[![4pLUWq.png](https://z3.ax1x.com/2021/09/12/4pLUWq.png)](https://imgtu.com/i/4pLUWq)

## 3.3 客户端流式 RPC（Client-side streaming RPC）

客户端流式 RPC，也是单向流，不过是由客户端发送流式数据罢了。

### 3.3.1 proto

其实关键就是在客户端发送的数据前数据前加上 `stream` 关键字

~~~protobuf
service ClientSide {
  //一个ClientSideHello的方法
  rpc ClientSideHello (stream ClientSideRequest) returns (ClientSideResp) {}
}
~~~

然后运行 `protoc --go_out=plugins=grpc:. *.proto` 生成对应的代码。

### 3.3.2 实现服务端代码

实现代码的话，跟客户端流模式代码大同小异。

#### 3.3.2.1 定义我们的服务

首先定义我们的服务 `ClientSideService` 并且实现`ClientSideHello`方法。

~~~go
type ClientSideService struct {
}

func (c *ClientSideService) ClientSideHello(server pb.ClientSide_ClientSideHelloServer) error {
	for i := 0; i < 5; i++ {
		recv, err := server.Recv()
		if err != nil {
			return err
		}
		log.Println("客户端信息：", recv)
	}
	//服务端最后一条消息发送
	err := server.SendAndClose(&pb.ClientSideResp{Message: "关闭"})
	if err != nil {
		return err
	}
	return nil
}

~~~

然后在 [server包](https://github.com/CodeFish-xiao/blogs/blob/main/gRPCAction/code/grpc-3/server/main.go) 中注册service

~~~go
    pb.RegisterClientSideServer(grpcServer, &services.ClientSideService{})
~~~~

#### 3.3.2.2 运行我们的服务

运行后：

[![49P4dx.png](https://z3.ax1x.com/2021/09/12/49P4dx.png)](https://imgtu.com/i/49P4dx)

### 3.3.2 实现客户端代码

代码如下：

~~~go
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
    for i := 0; i < 5; i++ {
    //通过 Send方法发送流信息 
	    err = res.Send(&pb.ClientSideRequest{Name: "客户端流式"})
	    if err != nil {
		    return
	    }
    }
    // 打印返回值 
    log.Println(res.CloseAndRecv())
}
~~~

### 3.3.3 客户端流模式运行样例

编写完客户端代码后，运行可见：客户端发送流请求，之后服务端进行打印，5次后服务端发送关闭流信息，客户端收到关闭信息，并且关闭了流：


客户端：

[![49POOA.png](https://z3.ax1x.com/2021/09/12/49POOA.png)](https://imgtu.com/i/49POOA)

服务端：

[![49PxTP.png](https://z3.ax1x.com/2021/09/12/49PxTP.png)](https://imgtu.com/i/49PxTP)


## 3.4 双向流式 RPC（Bidirectional streaming RPC）

客户端和服务端双方使用读写流去发送一个消息序列，两个流独立操作，双方可以同时发送和同时接收。

### 3.4.1 proto

在请求值和返回值前加上 `stream` 关键字

~~~protobuf
service Bidirectional {
  //一个BidirectionalHello的方法
  rpc BidirectionalHello (stream BidirectionalRequest) returns (stream BidirectionalResp) {}
}
~~~

然后运行 `protoc --go_out=plugins=grpc:. *.proto` 生成对应的代码。

### 3.4.2 实现服务端代码

#### 3.4.2.1 定义我们的服务

首先定义我们的服务 `BidirectionalService` 并且实现`BidirectionalHello`方法。

~~~go
type BidirectionalService struct {
}

func (b *BidirectionalService) BidirectionalHello(server pb.Bidirectional_BidirectionalHelloServer) error {
    defer func() {
        log.Println("客户端断开链接")
    }()
    for  {
    	//获取客户端信息 
    	recv, err := server.Recv()
    	if err != nil {
    		return err
    	}
    	log.Println(recv) 
    	//发送服务端信息 
    	err = server.Send(&pb.BidirectionalResp{Message: "服务端信息"})
    	if err != nil {
    		return err
    	}
    }
}

~~~

然后在 [server包](https://github.com/CodeFish-xiao/blogs/blob/main/gRPCAction/code/grpc-3/server/main.go) 中注册service

~~~go
	pb.RegisterBidirectionalServer(grpcServer, &services.BidirectionalService{})
~~~~

#### 3.4.2.2 运行我们的服务


运行后：
[![49P4dx.png](https://z3.ax1x.com/2021/09/12/49P4dx.png)](https://imgtu.com/i/49P4dx)

### 3.4.2 实现客户端代码

代码如下：

~~~go
func Bidirectional() {
	// 连接服务器 
	conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close() 
	// 建立gRPC连接 
	grpcClient := pb.NewBidirectionalClient(conn) 
	//获取流信息 
	stream, err := grpcClient.BidirectionalHello(context.Background())
	if err != nil {
		log.Fatalf("get BidirectionalHello stream err: %v", err)
	}
	
	for n := 0; n < 5; n++ {
		err := stream.Send(&pb.BidirectionalRequest{Name: "双向流 rpc " + strconv.Itoa(n)})
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
~~~
双向流模式，客户端需要从服务端获取流链接，之后双方都可以通过该流进行传输

### 3.4.3 双向流模式运行样例
因为grpc处理了断开链接后的处理，所以在客户端断开后，defer的代码可以运行并且输出信息。

客户端：
[![498DTf.png](https://z3.ax1x.com/2021/09/12/498DTf.png)](https://imgtu.com/i/498DTf)

服务端：
[![4986fg.png](https://z3.ax1x.com/2021/09/12/4986fg.png)](https://imgtu.com/i/4986fg)
### 3.5 小结


简单模式在上一节已经有说过，这次将其他几个交互模式都阐述了一遍，基本对大部分业务场景都够用了。但是在实际开发中，我们更多会需要很多东西：超时控制，负载均衡，权限控制，数据验证等功能，后续将会慢慢道来。
