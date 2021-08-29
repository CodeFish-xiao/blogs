package services

import (
	"github.com/CodeFish-xiao/blogs/gRPCAction/code/grpc-3/pb"
	"log"
)

type ServerSideService struct {
}

func (s *ServerSideService) ServerSideHello(request *pb.ServerSideRequest, server pb.ServerSide_ServerSideHelloServer) error {
	log.Println(request.Name)
	for n := 0; n < 5; n++ {
		// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes
		err := server.Send(&pb.ServerSideResp{Message: "你好"})
		if err != nil {
			return err
		}
	}
	return nil
}
