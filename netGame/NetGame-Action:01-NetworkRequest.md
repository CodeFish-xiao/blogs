# 网络请求

在日常的游戏开发或者软件开发中，大家基本都会使用到或者了解到Http请求或者Socket请求。最常见就是http请求。只要你是C/S架构或者B/S架构的软件，你基本上就会接触到网络请求。

在游戏软件中基本会将游戏场景分为：强联网场景和弱联网场景。我们下面以这两个场景进行展开讲述所用到的技术点以及该如何使用他们。

## 联网场景简介

### 弱联网场景

弱联网场景基本上是对于游戏状态而言的。大部分单机游戏基本上都是弱联网游戏。

对于这类游戏而言，联网的场景一般就是属于登录，购买，存档。 需要什么状态或者信息就直接从服务器获取或者更改即可。通讯方式基本上为一来一回。

部分卡牌游戏也会采用弱联网的方式进行交互，一定时间轮询房间内的状态，然后在客户端表现出来。

这类场景基本上都会直接使用Http请求或者是简单的RPC请求进行交互。同时服务端也会尽量设计为无状态服务，方便横向扩展。

### 强联网场景

强联网游戏基本上你不连接网络就玩不了了，或者当你网络不好的时候得到的体验就相当的差。

一般玩家之间对战，聊天服务这类的游戏都是强联网的场景，即时性要求较高。并且并不是一来一回的请求模式，服务端可以主动推送信息给不同的客户端，不需要等待客户端主动查询状态。

这类场景基本都是直接使用基础通讯进行交互，视情况而言会使用TCP，UDP或者KCP，以及基于http升级的Websocket协议。


## 协议和基础应用

### [Http协议](https://developer.mozilla.org/zh-CN/docs/Web/HTTP)

大部分的地方都有Http的协议介绍，这了就不赘述了。

我们直接开门见山介绍一下客户端以及服务端的用法和代码样例。

#### Unity发起Http请求

从Unity 2017开始中基本会使用 [`UnityWebRequest`](https://docs.unity3d.com/cn/2017.4/Manual/UnityWebRequest.html)  进行网络请求 。

```C#
public class HttpDemo : MonoBehaviour
{
    public string url;
    
    // Start is called before the first frame update
    void Start()
    {
        url = "http://127.0.0.1:3001/";
        StartCoroutine(Get());
    }
    IEnumerator Get()
    {
        UnityWebRequest webRequest = new UnityWebRequest(url,UnityWebRequest.kHttpVerbGET);
        webRequest.downloadHandler = new DownloadHandlerBuffer();
        yield return webRequest.SendWebRequest();
        if(webRequest.isNetworkError || webRequest.isHttpError) {
            Debug.Log(webRequest.error);
        }
        // 或者以二进制数据格式检索结果
        byte[] results = webRequest.downloadHandler.data;
        
    }
}
```

我们来简单讲解下这段代码：

主要看 Get() 方法：

- 返回协程形式迭代器：   ` IEnumerator Get()`

- 实例化 UnityWebRequest 类，并且使用Url和Get方法进行初始化：

`UnityWebRequest webRequest = new UnityWebRequest(url,UnityWebRequest.kHttpVerbGET);` 

- 初始化 webRequest.downloadHandler ，不然返回后会为空。

`webRequest.downloadHandler = new DownloadHandlerBuffer();`

- 等待返回后判断是否有没有错误进行解析

#### Golang 提供Http服务器

使用Golang起一个Http服务器很简单：

```
func main() {
   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello")
   })
   s := &http.Server{
      Addr:           ":3001",
      ReadTimeout:    10 * time.Second,
      WriteTimeout:   10 * time.Second,
      MaxHeaderBytes: 1 << 20,
   }
   log.Fatal(s.ListenAndServe())
}
```

我们来简单讲解一下下面的代码：

- 注册路由方法：

```
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Hello")
})
```

- 实例化http.Server：

```
s := &http.Server{
   Addr:           ":3001",
   ReadTimeout:    10 * time.Second,
   WriteTimeout:   10 * time.Second,
   MaxHeaderBytes: 1 << 20,
}
```

- 启动Server监听：

```
log.Fatal(s.ListenAndServe())
```

### 长链接

在长链接里，一般会使用TCP，UDP或者KCP，以及基于http升级的Websocket协议。这几种协议的具体差别很多地方都有。

这里以TCP为例子讲解如何进行TCP链接的使用。

#### Echo 程序

Echo程序是网络编程中最基础的案例。建立网络连接后，客户端向服务端发送一行文本，服务端收到后将文本发送回客户端，其实就是最简单版的一来一回的请求代码：

我们需要自己实现以下几步：

1. 建立链接
2. 发送消息
3. 接受并且回显消息
4. 关闭链接

#### Unity 长链接实现

##### 搭建UI面板

这里我们会第一次接触Unity的脚本编程。可以先简单跟着做就行，这里涉及到的东西在后续的章节会进行详细解说。



##### 脚本编写

在Unity中需要使用 `Net.Sockets` 包进行Sockets网络开发。

```
public class Echo:MonoBehaviour { 
    //定义套接字
    Socket socket;
    //UGUI 输入框
    public InputField InputFeld;
    //回显用的text组件
    public Text text;
    //请求的服务端 host
    public string host;
    //服务端端口
    public int port;
    //初始化值
    public void Start()
    {
        host = "127.0.0.1";
        port = 8972;
    }
    
    //点击连接按钮事件
    public void Connection()
    {
        //Socket
        socket = new Socket(AddressFamily.InterNetwork,SocketType.Stream, ProtocolType.Tcp);
        //Connect
        socket.Connect(host, port);
    }
    
    //点击发送按钮
    public void Send()
    {
        //Send 
        // 获取组件的text值
        string sendStr = InputFeld.text;
        byte[] sendBytes = System.Text.Encoding.Default.GetBytes(sendStr);
        //发送值到服务器
        socket.Send(sendBytes);
        //Recv
        byte[] readBuff = new byte[1024]; 
        //接收服务器返回值
        int count = socket.Receive(readBuff); 
        string recvStr = System.Text.Encoding.Default.GetString(readBuff, 0, count); 
        text.text = recvStr; 
        //Close 关闭连接
        socket.Close();
    }
} 
```

#### Golang 服务端代码

```
package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8972")
	if err != nil {
		panic(err)
	}
	var connections []net.Conn
	defer func() {
		for _, conn := range connections {
			conn.Close()
		}
	}()
	for {
		conn, e := ln.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Timeout() {
				log.Printf("accept temp err: %v", ne)
				continue
			}
			log.Printf("accept err: %v", e)
			return
		}
		go handleConn(conn)
		connections = append(connections, conn)
		if len(connections)%100 == 0 {
			log.Printf("total number of connections: %v", len(connections))
		}
	}
}

//处理链接
func handleConn(conn net.Conn) {
	c := NewClient(conn)
	go c.Echo()
}

// Client 客户端结构体
type Client struct {
	net.Conn //保存链接
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn}
}

// Echo Echo逻辑代码，收到消息直接发回客户端即可
func (c Client) Echo() {
	buf := make([]byte, 1024)
	for {
		c.Read(buf)
		c.Write(buf)
	}
}

```
