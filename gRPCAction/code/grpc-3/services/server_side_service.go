package services

import "github.com/CodeFish-xiao/blogs/gRPCAction/code/grpc-3/pb"

type ServerSideService struct {
}

func (s *ServerSideService) ServerSideHello(request *pb.ServerSideRequest, server pb.ServerSide_ServerSideHelloServer) error {
	panic("implement me")
}
