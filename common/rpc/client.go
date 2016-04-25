package rpc

import (
	log "github.com/Sirupsen/logrus"
	"google.golang.org/grpc"
)

func Connect(name string) (*grpc.ClientConn, error) {
	conn := new(grpc.ClientConn)
	config := NewConfig(name)
	address, err := config.ServerAddress()
	if err != nil {
		return conn, err
	}
	log.Printf("Service: %s,Address: %s", name, address)
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return conn, err
	}
	return conn, nil
}
