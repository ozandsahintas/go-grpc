// +acceptance

package test

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"odesch.com/odesch/grpc-app/protos/rocket/v1"
)

func GetClient() rocket.RocketServiceClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	rocketClient := rocket.NewRocketServiceClient(conn)
	return rocketClient
}
