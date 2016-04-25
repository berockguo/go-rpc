package server

import (
	pb "gitlab.ucloudadmin.com/ucre/rpc/pkg/ucrehelloworld/proto"
	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	resp := new(pb.SayHelloResponse)
	str := "hello," + *in.Name
	resp.Ret.Retmsg = &str
	return resp, nil
}
