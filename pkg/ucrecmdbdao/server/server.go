package server

import (
	pb "gitlab.ucloudadmin.com/ucre/rpc/pkg/ucrehelloworld/protos"
	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{Message: "Hello " + in.Name}, nil
}
