package client

import (
	log "github.com/Sirupsen/logrus"

	"gitlab.ucloudadmin.com/ucre/rpc/common/rpc"
	"gitlab.ucloudadmin.com/ucre/rpc/ucrehelloworld/pkg/define"
	pb "gitlab.ucloudadmin.com/ucre/rpc/ucrehelloworld/protos"
	"golang.org/x/net/context"
)

func SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	resp := new(pb.SayHelloResponse)
	conn, err := rpc.Connect(define.ServiceName)
	if err != nil {
		return resp, err
	}
	defer conn.Close()
	c := pb.NewUcreHelloWorldClient(conn)
	resp, err = c.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return resp, err
	}
	log.Printf("Ucre SayHello: %s", resp.Message)
	return resp, nil
}
