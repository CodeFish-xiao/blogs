# gRPC应用实战：（番外）实现原生的RPC调用
## 99.1 前言
我们在掌握gRPC的时候，一定要简单知道整个调用流程是怎么样的。在第一节有简单的说过，golang自己有原生的RPC调用，那么，我们今天就来了解一下Golang原生的RPC调用以及调用的部分流程。

## 99.2 实现原生RPC
### 99.2.1 定制交互消息 
正像我们使用gRPC一样，在gRPC中需要先编写pb文件确定接口交互的消息体，那么在我们原生的也是一样的：
~~~go
type HelloRequest struct {
	Name string
}
~~~

简单一点，我们只需要定义一个请求的消息体就行，返回体的话就用字符串就好，等后面就知道了。
### 99.2.2 客户端代码编写
代码如下：
~~~go
func main() {
	//获取服务端地址
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("DialHTTP err :", err)
	}
	//进行请求结构体的声明
	req := &api.HelloRequest{Name: "codefish"}
	var resp string
	//同步调用
	err = client.Call("HelloService.Hello", req, &resp)
	log.Println(resp)
	//异步调用
	call := client.Go("HelloService.HelloGO", req, &resp, nil)
	_ = <-call.Done
	log.Println(resp)
}
~~~
简单解释一下：
- 首先需要进行对于服务端地址的获取，并且建立链接。
- 在golang中原生实现了两种请求方式：
  - 同步调用方式，调用HelloService的Hello方法
  - 异步调用方式，调用HelloService的HelloGO方法
- 之后就是看服务端的信息返回或者是处理了

### 99.2.3服务端代码实现
服务端代码处理：
老规矩，首先实现service层：
~~~go
type Service interface {
//Hello a and b
Hello(req api.HelloRequest, ret *string) error

//HelloGO a and b
HelloGO(req api.HelloRequest, ret *string) error
}
type HelloService struct {
}

func (s *HelloService) Hello(req api.HelloRequest, ret *string) error {
log.Println("接客了：", req.Name)
*ret = "谢谢大爷"
return nil
}
func (s *HelloService) HelloGO(req api.HelloRequest, ret *string) error {
log.Println("接客了：", req.Name)
*ret = "异步，谢谢大爷"
return nil
}

// NewService 构造函数
func NewService() Service {
return &HelloService{}
}

~~~
这里的返回直接就是一串字符串，做简单演示，