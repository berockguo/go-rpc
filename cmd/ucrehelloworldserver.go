package main

import (
	"gitlab.ucloudadmin.com/ucre/rpc/pkg/common/rpc"
	"gitlab.ucloudadmin.com/ucre/rpc/pkg/ucrehelloworld/define"
	pb "gitlab.ucloudadmin.com/ucre/rpc/pkg/ucrehelloworld/protos"
	"gitlab.ucloudadmin.com/ucre/rpc/pkg/ucrehelloworld/server"
)

func main() {
	s, err := rpc.NewServer(define.ServiceName)
	if err != nil {
		return
	}

	pb.RegisterUcreHelloWorldServer(s.GRpcServer, &server.Server{})
	err = s.Serve()
	if err != nil {
		return
	}
}
