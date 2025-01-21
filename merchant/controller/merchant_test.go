package controller

import (
	"context"
	pb "merchant/pb/merchantpb"
	"merchant/repo"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/metadata"
)

var merchantRepoMock = &repo.MerchantRepositoryMock{Mock: mock.Mock{}}
var merchantServiceMock = Merchant{Repository: merchantRepoMock}

func TestShowAllProductsNoValidation(t *testing.T) {
	token := ""
	md := metadata.Pairs("Authorization", token)
	ctxWithToken := metadata.NewOutgoingContext(context.Background(), md)

	_, err := (merchantServiceMock).ShowAllProducts(ctxWithToken, &pb.ShowAllProductRequest{})
	assert.NotNil(t, err)
}

func TestDeleteProduct(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVtYWlsQG1haWwuY29tIiwiZXhwIjoxNzM3NTQzMTM0LCJyb2xlIjoibWVyY2hhbnQiLCJ1c2VyX2lkIjoiMDE3YzBhNzQtYzgwYS00NDRiLWFjYmEtOGFhMWUwMGQ3YmNjIn0.6Idb7aBxfIT-nw3yfKuh0g-uylEYWn3ICIO8PgR0fIw"
	md := metadata.Pairs("Authorization", token)
	ctxWithToken := metadata.NewOutgoingContext(context.Background(), md)

	_, err := (merchantServiceMock).DeleteProduct(ctxWithToken, &pb.DeleteProductRequest{
		ProductId: "xxxxxx",
	})
	assert.NotNil(t, err)
}
