package main

import (
	log "github.com/Sirupsen/logrus"

	ucrehelloworldclient "gitlab.ucloudadmin.com/ucre/rpc/ucrehelloworld/pkg/client"
	pb "gitlab.ucloudadmin.com/ucre/rpc/ucrehelloworld/protos"
	"golang.org/x/net/context"
)

func main() {
	r, err := ucrehelloworldclient.SayHello(context.Background(), &pb.SayHelloRequest{Name: "TestName"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Ucre SayHello: %s", r.Message)
}
