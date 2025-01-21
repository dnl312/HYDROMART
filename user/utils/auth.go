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

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return claims, err
	}

	// Check if the token is valid
	if !token.Valid {
		return claims, errors.New("invalid token")
	}

	// Check token expiration
	// if time.Now().Unix() > claims.Exp {
	// 	return claims, errors.New("token has expired")
	// }

	return claims, nil
}

// func RecoverUser(tokenString string) (map[string]interface{}, error) {
// 	log.Printf("Received token2: %v", tokenString)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("JWT_SECRET")), nil
// 	})

// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse token: %v", err)
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		user := make(map[string]interface{})
// 		if userID, ok := claims["user_id"].(string); ok {
// 			user["user_id"] = userID
// 		} else {
// 			return nil, errors.New("invalid user_id in token claims")
// 		}
// 		if email, ok := claims["email"].(string); ok {
// 			user["email"] = email
// 		} else {
// 			return nil, errors.New("invalid email in token claims")
// 		}
// 		if role, ok := claims["role"].(string); ok {
// 			user["role"] = role
// 		} else {
// 			return nil, errors.New("invalid role in token claims")
// 		}
// 		log.Printf("User: %v", user)
// 		return user, nil
// 	}

// 	return nil, errors.New("invalid token")
// }
