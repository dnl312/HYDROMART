package utils

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestGetTokenStringFromContext(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVtYWlsQG1haWwuY29tIiwiZXhwIjoxNzM3NTQzMTM0LCJyb2xlIjoibWVyY2hhbnQiLCJ1c2VyX2lkIjoiMDE3YzBhNzQtYzgwYS00NDRiLWFjYmEtOGFhMWUwMGQ3YmNjIn0.6Idb7aBxfIT-nw3yfKuh0g-uylEYWn3ICIO8PgR0fIw"
	md := metadata.Pairs("Authorization", token)
	ctxWithToken := metadata.NewOutgoingContext(context.Background(), md)

	_, err := GetTokenStringFromContext(ctxWithToken)
	assert.NotNil(t, err)
}

func TestRecoverUserEmptyToken(t *testing.T) {
	token := ""
	_, err := RecoverUser(token)
	assert.Equal(t, err, errors.New("failed to parse token: token contains an invalid number of segments"))

}

func TestRecoverUserWrongToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVtYWlsQG1haWwuY29tIiwiZXhwIjoxNzM3NTM2MzQ2LCJyb2xlIjoibWVyY2hhbnQiLCJ1c2VyX2lkIjoiMDE3YzBhNzQtYzgwYS00NDRiLWFjYmEtOGFhMWUwMGQ3YmNjIn0.TjqUiYExuPd70zOsJVah"
	_, err := RecoverUser(token)
	assert.Equal(t, err, errors.New("failed to parse token: signature is invalid"))

}
