# gRPC应用实战：（二）gRPC环境搭建以及简单demo
## 2.1 gRPC环境安装
### 2.1.1 golang环境安装
接下来的所有的教程以及代码都是golang代码进行演示，所以我们需要安装golang环境：
下载地址：[golang 官网](https://golang.google.cn/dl/)

安装完在你的命令行工具输入
```
go version
```
就可以看到对应的下载版本

[![hSBTAg.png](https://z3.ax1x.com/2021/08/22/hSBTAg.png)](https://imgtu.com/i/hSBTAg)
然后就可以了。

### 2.1.2 protoc 安装
在 gRPC 开发中，我们常常需要与 Protobuf 进行打交道，而在编写了.proto 文件后，我们会需要到一个编译器，就是protoc。这个工具呢可以在GitHub上直接下载

[https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)

大家可以在这个网站自行下载安装（因为之前这里安装完了，也不想重装了，可能记得不太清了，你们直接试一试就好了，等下次重装再丰富这里）
一般操作是把protoc编译文件所在的目录配到环境变量里就好了
安装完后：

运行
>protoc --version

就可以看见版本信息了

[![hSy129.png](https://z3.ax1x.com/2021/08/22/hSy129.png)](https://imgtu.com/i/hSy129)


### 2.1.2 安装相关依赖包

安装 golang 的proto工具包

> go get -u github.com/golang/protobuf/proto

安装 goalng 的proto编译支持

>go get -u github.com/golang/protobuf/protoc-gen-go

安装 gRPC 包

>go get -u google.golang.org/grpc

这样子就基本上结束对于环境的安装了


## 2.2 gRPC简单demo
### 2.2.1 protocol buffer 语法

在gRPC中主要以protocol buffer来定义api以及服务，所以我们需要先了解一下protocol buffer的语法。protocol buffer主要使用中有两个版本：proto2和proto3，这里呢，推荐大家使用proto3进行日常开发。

首先：一切的学习都要学会看文档：这里是谷歌对于protocol buffer proto3版本的的文档[Language Guide (proto3) ](https://developers.google.com/protocol-buffers/docs/proto3#generating)

当然你要是觉得英文看不过去的话，这里还有鸟窝大大转发的千念飞羽大大翻译的[Protobuf3语言指南](https://blog.csdn.net/u011518120/article/details/54604615)可以让大家学习。

接下来是一个最最最简单protocol buffer demo，可搭配注释食用：

~~~protobuf
//指定proto3语法
syntax = "proto3";
//包名
package helloworld;

//一个为Hello的服务（可定义多个服务,每个服务可定义多个方法）
service Hello {
    //一个SayHello的方法
    rpc SayHello (HelloRequest) returns (HelloResp) {}
}
// 定义发送请求信息
message HelloRequest {
     // 定义发送的参数
    // 参数类型 参数名 标识号(不可重复)
    string name = 1;
}
// 定义响应信息
message HelloResp {
    string message = 1;
}
~~~

### 2.2.2 编译proto文件

在编译文件前，我们先搭建一个最简单的项目目录，用来具体演示相关操作：

[![hSRQXj.png](https://z3.ax1x.com/2021/08/22/hSRQXj.png)](https://imgtu.com/i/hSRQXj)

client：gRPC客户端代码
pb：存放公共pb文件以及编译文件
server：gRPC服务端代码


接下来可以编译编写好的代码

通过下面的命令就可以直接生成对应语言的代码，具体代码可见：[这里](https://github.com/CodeFish-xiao/blogs/tree/main/gRPCAction/code/grpc-2)
~~~
protoc --go_out=plugins=grpc:. *.proto
~~~

### 2.2.3 编写服务端代码

接下来便可以编写服务端代码（结合注释食用即可）

~~~go
import (
	"context"
	pb "github.com/CodeFish-xiao/blogs/gRPCAction/code/grpc-2/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)
// HelloService 定义我们的服务
type HelloService struct {
	
}

// SayHello 实现SayHello方法
func (s *HelloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResp, error)  {
	log.Println(req.Name)
	return &pb.HelloResp{Message: "hello ,I'm codefish "},nil
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
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	pb.RegisterHelloServer(grpcServer,&HelloService{})
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Panic("grpcServer.Serve err: %v", err)
	}
}
~~~

运行客户端可见：

[![hSIfv8.png](https://z3.ax1x.com/2021/08/22/hSIfv8.png)](https://imgtu.com/i/hSIfv8)

这就是运行成功了
### 2.2.4 编写客户端代码

接下来编写客户端代码：结合注释食用：

~~~go
import (
	"context"
	pb "github.com/CodeFish-xiao/blogs/gRPCAction/code/grpc-2/pb"
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
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := grpcClient.SayHello(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}
~~~

运行后可见：

[![hSo3xf.png](https://z3.ax1x.com/2021/08/22/hSo3xf.png)](https://imgtu.com/i/hSo3xf)

并且在server的控制台可以看到打印的参数：

[![hSo0Gq.png](https://z3.ax1x.com/2021/08/22/hSo0Gq.png)](https://imgtu.com/i/hSo0Gq)

## 总结

这样子便可以搭建一个简单的gRPC服务了，可以让客户端和服务端进行简单的业务交互了，当然这才刚刚开始。。。


本章代码可见[codefish的github](https://github.com/CodeFish-xiao/blogs/tree/main/gRPCAction/code/grpc-2)

