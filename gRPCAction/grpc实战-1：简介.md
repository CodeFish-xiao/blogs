# gRPC应用实战：（一）

我们将介绍 Go 语言中最流行的 RPC 框架：gRPC，并带你探索其相对应的技术栈。

首先我们需要对RPC有所了解，然后再对他们的传输协议进行了解并且建立对应的demo，然后一步步深入并且完成较为完善的搭建体系。

## 1.1 RPC

### 1.1.1 什么是RPC？

RPC 代指远程过程调用（Remote Procedure Call）就是允许程序调用另一个地址空间（通常是共享网络的另一台机器上）的过程或函数，且不需要显式编码这个远程调用的细节。

通俗来说就是调用其他进程内的方法，并且获得对应的返回值。

### 1.1.2 RPC包括什么？

整个RPC调用基本上会包括以下3点：

1. 通讯协议
2. 寻址
3. 数据序列化

RPC 框架的目标就是让远程服务调用更加简单、透明。一个较为成熟的RPC 框架负责屏蔽底层的传输方式（TCP 或者 UDP）、序列化方式（XML/Json/
二进制）和通信细节。服务调用者可以像调用本地接口一样调用远程的服务提供者，而不需要关心底层通信细节和调用过程。

### 1.1.3 业界主流的 RPC 框架

业界主流的 RPC 框架整体上分为三类：

1. 支持多语言的 RPC 框架，比较成熟的有 Google 的 gRPC、Apache（Facebook）的 Thrift；
2. 只支持特定语言的 RPC 框架，例如新浪微博的 Motan；
3. 支持服务治理等服务化特性的分布式服务框架，其底层内核仍然是 RPC 框架, 例如阿里的 Dubbo。

## 1.2 序列化协议

常见的序列化协议有XML,jSON,Protobuf，一般来说用于RPC调用时基本就是json跟Protobuf两种协议的事情，所以这里只谈这两种协议：

### 1.2.1 Protobuf

#### 1.2.1.1 什么是Protobuf

Protocol Buffers（Protobuf）是一种与语言、平台无关，可扩展的序列化结构化数据的数据描述语言，我们常常称其为 IDL，常用于通信协议，数据存储等等，相较于
JSON、XML，它更小、更快，因此也更受开发人员的青眯。在gRPC中就是使用的Protobuf作为传输协议。

#### 1.2.1.2 Protobuf的基本语法

~~~protobuf
syntax = "proto3";

package helloworld;

service Hello {
    rpc SayHello (HelloRequest) returns (Resp) {}
}

message HelloRequest {
    string name = 1;
}

message HelloResp {
    string message = 1;
}
~~~

1. 第一行（非空的非注释行）声明使用 proto3 语法。如果不声明，将默认使用 proto2 语法。
2. 定义名为 Hello 的 RPC 服务（Service),中有SayHello的方法，以及有入参HelloRequest，和返回值HelloResp
3. 定义 HelloRequest、HelloResp 消息体。

在编写完.proto 文件后，我们一般会进行编译和生成对应语言的 proto 文件操作，这个时候 Protobuf 的编译器会根据选择的语言不同、调用的插件情况，生成相应语言的 Service Interface Code 和 Stubs

### 1.2.2 Json跟Protobuf对比

Json比较常见就不详细介绍了，直接两者进行对比即可：

|     | Json  |Protobuf|
|  ----  | ----  | ---- |
| 传输形式|字符串  | 二进制|
| 可读性  | 好       |不好 |
| 解析速度  |较慢  |较快 |
|约定性   | 较强 |很强 |

## 1.3 gRPC

### 1.3.1 什么是 gRPC

gRPC 是一个高性能、开源和通用的 RPC 框架，面向服务端和移动端，基于 HTTP/2 设计。目前提供 C、Java 和 Go 语言等等版本，分别是：grpc、grpc-java、grpc-go，其中 C 版本支持
C、C++、Node.js、Python、Ruby、Objective-C、PHP 和 C# 支持

gRPC 的调用示例如下所示：
![](https://static001.geekbang.org/resource/image/6d/d9/6d9a335ad96491e4d610a31b5089a2d9.png)

一个Ruby客户端和安卓客户端都可以通过gRPC调用C++服务端

### 1.3.2 gRPC 特点

1. 语言中立，支持多种语言；
2. 基于 IDL 文件定义服务，通过 proto3 工具生成指定语言的数据结构、服务端接口以及客户端 Stub；
3. 通信协议基于标准的 HTTP/2 设计，支持双向流、消息头压缩、单 TCP 的多路复用、服务端推送等特性，这些特性使得 gRPC 在移动端设备上更加省电和节省网络流量；
4. 序列化支持 PB（Protocol Buffer）和 JSON，PB 是一种语言无关的高性能序列化框架，基于 HTTP/2 + PB, 保障了 RPC 调用的高性能。

## 1.4 小结

简单介绍了一下RPC以及序列化协议还有gprc，下次就是开始简单搭建demo了。