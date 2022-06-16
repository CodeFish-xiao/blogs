package decorator

import (
	"fmt"
	"testing"
)

func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d \n", a, b, c)
	return a + b + c
}

func bar(a, b string) string {
	fmt.Printf("%s, %s \n", a, b)
	return a + b
}
func TestInterfaceDecorator(t *testing.T) {

	type MyFoo func(int, int, int) int
	var myfoo MyFoo
	err := InterfaceDecorator(&myfoo, foo)
	if err != nil {
		fmt.Println(err)
		return
	}
	myfoo(1, 2, 3)
	//不申明函数的情况
	mybar := bar
	err = InterfaceDecorator(&mybar, bar)
	if err != nil {
		fmt.Println(err)
		return
	}
	mybar("hello,", "world!")
}
