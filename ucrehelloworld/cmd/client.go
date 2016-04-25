package main

import (
	"os"

	log "github.com/Sirupsen/logrus"

	"gitlab.ucloudadmin.com/ucre/rpc/common/rpc"
	pb "gitlab.ucloudadmin.com/ucre/rpc/ucrehelloworld/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	defaultName = "world"
	ServiceName = "ucrehelloworld"
)

func main() {
	// Set up a connection to the server.
	config := rpc.NewConfig(ServiceName)
	address, err := config.ServerAddress()
	if err != nil {
		log.Fatalf("read config file failed")
		return
	}
	log.Printf("Address: %s", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUcreHelloWorldClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.SayHelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Ucre SayHello: %s", r.Message)
}
