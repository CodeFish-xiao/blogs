package main

import (
	"github.com/CodeFish-xiao/blogs/gRPCAction/code/fish-rpc/server/service"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	stringService := service.NewService()
	err := rpc.Register(stringService)
	if err != nil {
		log.Fatal("Register err: ", err)
		return
	}
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen err: ", err)
	}
	err = http.Serve(l, nil)
	if err != nil {
		return
	}
}
