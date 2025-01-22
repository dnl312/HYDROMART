package utils

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"user/model"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func GetTokenStringFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("No metadata found")
		return "", status.Errorf(codes.Unauthenticated, "Unauthorized")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		log.Println("No authorization found")
		return "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	parts := strings.Split(authHeader[0], " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", status.Errorf(codes.Unauthenticated, "invalid authorization token format")
	}

	tokenString := parts[1]
	return tokenString, nil
}

func RecoverUser(tokenString string) (*model.Claims, error) {
	claims := &model.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return claims, err
	}

	if !token.Valid {
		return claims, errors.New("invalid token")
	}

	return claims, nil
}
