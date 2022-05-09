# gRPC应用实战：（三）gRPC拦截器

## 4.1 前言

在我们使用http请求时候，会有请求中间件，例如处理token，记录日志呀。所以呢，我们在使用gRPC的时候也会有这样的需求吧？
那么，这一节我们就来看看gRPC拦截器的写法。


在gRPC中呢，拦截器可以分为：
- 普通方法：一元拦截器（grpc.UnaryInterceptor）
- 流方法：流拦截器（grpc.StreamInterceptor）
## 4.2 一元拦截器
