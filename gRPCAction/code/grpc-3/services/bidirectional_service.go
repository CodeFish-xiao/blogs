package services

import "github.com/CodeFish-xiao/blogs/gRPCAction/code/grpc-3/pb"

type BidirectionalService struct {
}

func (b *BidirectionalService) ClientSideHello(server pb.ClientSide_ClientSideHelloServer) error {
	panic("implement me")
}
