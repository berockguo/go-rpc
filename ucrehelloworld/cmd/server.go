package main

import (
	"gitlab.ucloudadmin.com/ucre/rpc/common/rpc"
	pb "gitlab.ucloudadmin.com/ucre/rpc/ucrehelloworld/protos"
	"golang.org/x/net/context"
)

const (
	ServiceName = "ucrehelloworld"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{Message: "Hello " + in.Name}, nil
}

func main() {
	s, err := rpc.NewServer(ServiceName)
	if err != nil {
		return
	}
	pb.RegisterUcreHelloWorldServer(s.GRpcServer, &server{})
	err = s.Serve()
	if err != nil {
		return
	}
}
