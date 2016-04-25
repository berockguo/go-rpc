package main

import (
	"gitlab.ucloudadmin.com/ucre/rpc/common/rpc"
	"gitlab.ucloudadmin.com/ucre/rpc/ucrehelloworld/pkg/define"
	pb "gitlab.ucloudadmin.com/ucre/rpc/ucrehelloworld/protos"
	"golang.org/x/net/context"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{Message: "Hello " + in.Name}, nil
}

func main() {
	s, err := rpc.NewServer(define.ServiceName)
	if err != nil {
		return
	}
	pb.RegisterUcreHelloWorldServer(s.GRpcServer, &server{})
	err = s.Serve()
	if err != nil {
		return
	}
}
