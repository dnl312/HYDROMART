package config

import (
	"client/pb"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitAuthServiceClient() (*grpc.ClientConn, pb.AuthServiceClient) {
	conn, err := grpc.Dial(os.Getenv("AUTH_SERVICE_URI"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return conn, pb.NewAuthServiceClient(conn)
}