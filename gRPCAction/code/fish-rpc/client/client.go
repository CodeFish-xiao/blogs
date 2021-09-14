package main

import (
	"github.com/CodeFish-xiao/blogs/gRPCAction/code/fish-rpc/api"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("DialHTTP err :", err)
	}
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
