package rpc

import (
	"net"

	"gitlab.ucloudadmin.com/ucre/rpc/common/util"

	"google.golang.org/grpc"
)

type Server struct {
	ServiceName string
	GRpcServer  *grpc.Server
	Config      *Config
	Listener    net.Listener
}

func NewServer(name string) (*Server, error) {
	s := &Server{ServiceName: name}
	s.GRpcServer = grpc.NewServer()
	s.Config = NewConfig(name)
	port, err := s.Config.ServerPort()
	if err != nil {
		return s, err
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return s, &util.Error{"failed to listen" + port}
	}
	s.Listener = lis
	return s, nil
}
func (s *Server) Serve() error {
	err := s.GRpcServer.Serve(s.Listener)
	return err
}
