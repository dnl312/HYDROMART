package config

import (
	"auth/middleware"
	"auth/pb"
	"log"
	"net"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
)

func ListenAndServeGrpc(controller pb.AuthServiceServer) {
	port := os.Getenv("GRPC_PORT")
	
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatal(err)
	}
	
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(middleware.NewInterceptorLogger()),
		),
	)
	pb.RegisterAuthServiceServer(grpcServer, controller)

	log.Println("\033[36mGRPC server is running on port:", port, "\033[0m")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to server gRPC:", err)
	}
}
