package main

import (
	"net"

	log "github.com/Sirupsen/logrus"
	"gitlab.ucloudadmin.com/ucre/rpc/ucrehelloworld/pkg/config"
	pb "gitlab.ucloudadmin.com/ucre/rpc/ucrehelloworld/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{Message: "Hello " + in.Name}, nil
}

func main() {
	port, ok := config.ServerPort()
	if !ok {
		log.Fatalf("read config file failed")
		return
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUcreHelloWorldServer(s, &server{})
	s.Serve(lis)
}
