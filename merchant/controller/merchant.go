package controller

import (
	//pb "merchant/pb"
	"merchant/repo"
)

type Merchant struct {
	//pb.UnimplementedMerchantServiceServer
	Repository repo.MerchantInterface
}

func NewMerchantController(r repo.MerchantInterface) Merchant {
	return Merchant{
		Repository: r,
	}
}
