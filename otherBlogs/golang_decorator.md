# Go 语言的 修饰器模式（Decorator）

让我们从Hello World 开始看装饰器模式：

```go
package decorator

import "fmt"

func decorator(f func(s string)) func(s string) {
   return func(s string) {
      fmt.Println("Started")
      f(s)
      fmt.Println("Done")
   }
}

func Hello(s string) {
   fmt.Println(s)
}
```

```go
func TestDecorator(t *testing.T) {
   hello := decorator(Hello)
   hello("Hello")
}
```

可以看到：我们在调用decorator函数时，传入Hello函数，然后就会在函数上下文进行插入输出