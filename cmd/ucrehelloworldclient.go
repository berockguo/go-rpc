package main

import (
	ucrehelloworldclient "gitlab.ucloudadmin.com/ucre/rpc/pkg/ucrehelloworld/client"
	pb "gitlab.ucloudadmin.com/ucre/rpc/pkg/ucrehelloworld/proto"
	"golang.org/x/net/context"
)

func main() {
	name := "Ucre"
	ucrehelloworldclient.SayHello(context.Background(), &pb.SayHelloRequest{Name: &name})
}
