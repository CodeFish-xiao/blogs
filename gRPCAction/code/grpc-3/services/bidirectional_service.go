package services

import (
	"github.com/CodeFish-xiao/blogs/gRPCAction/code/grpc-3/pb"
	"log"
)

type BidirectionalService struct {
}

func (b *BidirectionalService) BidirectionalHello(server pb.Bidirectional_BidirectionalHelloServer) error {
	defer func() {
		log.Println("客户端断开链接")
	}()
	for {
		//获取客户端信息
		recv, err := server.Recv()
		if err != nil {
			return err
		}
		log.Println(recv)
		//发送服务端信息
		err = server.Send(&pb.BidirectionalResp{Message: "服务端信息"})
		if err != nil {
			return err
		}
	}
}
