package config

import (
	pb "client/pb/merchantpb"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitMerchantServiceClient() (*grpc.ClientConn, pb.MerchantServiceClient) {
	conn, err := grpc.Dial(os.Getenv("MERCHANT_SERVICE_URI"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return conn, pb.NewMerchantServiceClient(conn)
}
