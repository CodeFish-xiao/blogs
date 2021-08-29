package services

import "github.com/CodeFish-xiao/blogs/gRPCAction/code/grpc-3/pb"

type ClientSideService struct {
}

func (c *ClientSideService) BidirectionalHello(server pb.Bidirectional_BidirectionalHelloServer) error {
	panic("implement me")
}
