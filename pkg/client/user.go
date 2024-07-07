package client

import (
	pb "echoproject/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func ProvideUserClient() pb.UserServiceClient {
	conn, err := grpc.Dial(os.Getenv("GRPC_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return pb.NewUserServiceClient(conn)
}
