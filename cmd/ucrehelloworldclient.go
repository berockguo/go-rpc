package main

import (
	ucrehelloworldclient "gitlab.ucloudadmin.com/ucre/rpc/pkg/ucrehelloworld/client"
	pb "gitlab.ucloudadmin.com/ucre/rpc/pkg/ucrehelloworld/protos"
	"golang.org/x/net/context"
)

func main() {
	ucrehelloworldclient.SayHello(context.Background(), &pb.SayHelloRequest{Name: "TestName"})
}
