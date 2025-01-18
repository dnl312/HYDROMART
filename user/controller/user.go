package controller

import (
	//pb "user/pb"
	"user/repo"
)

type User struct {
	//pb.UnimplementedMerchantServiceServer
	Repository repo.UserInterface
}

func NewMerchantController(r repo.UserInterface) User {
	return User{
		Repository: r,
	}
}
