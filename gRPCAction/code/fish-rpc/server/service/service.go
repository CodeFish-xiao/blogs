package service

import (
	"github.com/CodeFish-xiao/blogs/gRPCAction/code/fish-rpc/api"
	"log"
)

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
