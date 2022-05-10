# gRPC应用实战：（三）gRPC拦截器

## 4.1 前言

在我们使用http请求时候，会有请求中间件，例如处理token，记录日志呀。所以呢，我们在使用gRPC的时候也会有这样的需求吧？
那么，这一节我们就来看看gRPC拦截器的写法。以及在生产环境下应该怎么使用以及编写拦截器。


在gRPC中呢，拦截器可以分为：
- 普通方法：一元拦截器（grpc.UnaryInterceptor）
- 流方法：流拦截器（grpc.StreamInterceptor）

## 4.2 一元拦截器

一元拦截器其实使用起来比较简单，让我们简单看看代码就知道了。

### 4.2.1 Server 端

验证的代码对于本质来说，其实就是对于metadata数据的处理验证而已，通过获取metadata信息，然后对应处理就好。

Check方法获取对应的字段，并且做校验：

```go
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
```

```go
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
opts := []grpc.ServerOption{}
opts = append(opts, grpc.UnaryInterceptor(checkInterceptor))
grpcServer := grpc.NewServer(opts...)
```

实现：

```go
func UnaryInterceptor(i UnaryServerInterceptor) ServerOption {
   return newFuncServerOption(func(o *serverOptions) {
      if o.unaryInt != nil {
         panic("The unary server interceptor was already set and may not be reset.")
      }
      o.unaryInt = i
   })
}
```



