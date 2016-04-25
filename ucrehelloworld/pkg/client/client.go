package client

import (
	log "github.com/Sirupsen/logrus"
	"google.golang.org/grpc"

	"gitlab.ucloudadmin.com/ucre/rpc/common/rpc"
	"gitlab.ucloudadmin.com/ucre/rpc/ucrehelloworld/pkg/define"
	pb "gitlab.ucloudadmin.com/ucre/rpc/ucrehelloworld/protos"
	"golang.org/x/net/context"
)

func Client() (*grpc.ClientConn, pb.UcreHelloWorldClient, error) {
	client := *new(pb.UcreHelloWorldClient)
	conn, err := rpc.Connect(define.ServiceName)
	if err != nil {
		return conn, client, err
	}
	client = pb.NewUcreHelloWorldClient(conn)
	return conn, client, nil
}

func SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	resp := new(pb.SayHelloResponse)
	conn, client, err := Client()
	defer conn.Close()
	if err != nil {
		return resp, err
	}
	resp, err = client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return resp, err
	}
	log.Printf("Ucre SayHello: %s", resp.Message)
	return resp, nil
}
