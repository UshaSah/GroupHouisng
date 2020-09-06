package sg_client

import (
	"log"

	pb "libraries/pb/sg"

	grpc "google.golang.org/grpc"
)

// TODO: Move to central constant config
const (
	address = "172.18.0.5:50051"
)

type SGClient struct {
	Client *pb.SGClient
}

func (sgc *SGClient) Dial() *grpc.ClientConn {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}