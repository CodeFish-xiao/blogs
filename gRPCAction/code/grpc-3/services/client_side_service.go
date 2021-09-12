package services

import (
	"github.com/CodeFish-xiao/blogs/gRPCAction/code/grpc-3/pb"
	"log"
)

type ClientSideService struct {
}

func (c *ClientSideService) ClientSideHello(server pb.ClientSide_ClientSideHelloServer) error {

	for i := 0; i < 5; i++ {
		recv, err := server.Recv()
		if err != nil {
			return err
		}
		log.Println("客户端信息：", recv)
	}
	//服务端最后一条消息发送
	err := server.SendAndClose(&pb.ClientSideResp{Message: "关闭"})
	if err != nil {
		return err
	}
	return nil
}
