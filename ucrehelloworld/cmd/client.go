package main

import (
	ucrehelloworldclient "gitlab.ucloudadmin.com/ucre/rpc/ucrehelloworld/pkg/client"
	pb "gitlab.ucloudadmin.com/ucre/rpc/ucrehelloworld/protos"
	"golang.org/x/net/context"
)

func main() {
	ucrehelloworldclient.SayHello(context.Background(), &pb.SayHelloRequest{Name: "TestName"})
}
