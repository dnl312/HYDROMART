package config

import (
	pb "client/pb/userpb"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitOrderServiceClient() (*grpc.ClientConn, pb.OrderClient) {
	conn, err := grpc.Dial(os.Getenv("USER_SERVICE_URI"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return conn, pb.NewOrderClient(conn)
}
