package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"merchant/model"
	"net/http"
	"os"
	"strings"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func JWTAuth(ctx context.Context) (context.Context, error) {
	log.Println("JWTAuth middleware")
	tokenString, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		fmt.Println(tokenString)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok {
		return ctx, nil
	}
	return nil, status.Error(codes.Unauthenticated, "failed to verify jwt")
}

func CustomJWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization Header")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization Header Format")
		}
		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		}

		c.Set("user", token)
		return next(c)
	}
}

func GetUserIDFromToken(c echo.Context) (model.User, error) {

	userToken := c.Get("User")
	if userToken == nil {
		return model.User{}, errors.New("user token is missing from context")
	}

	token, ok := userToken.(*jwt.Token)
	if !ok {
		return model.User{}, errors.New("invalid token format")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return model.User{}, errors.New("invalid or malformed token claims")
	}

	// Extract userID
	user := model.User{
		UserID: claims["user_id"].(string),
		Role:   claims["role"].(string),
		Email:  claims["email"].(string),
	}

	_, ok = claims["user_id"].(string)
	if !ok {
		return model.User{}, errors.New("user ID not found or invalid in token claims")
	}
	_, ok = claims["role"].(string)
	if !ok {
		return model.User{}, errors.New("role not found or invalid in token claims")
	}
	_, ok = claims["email"].(string)
	if !ok {
		return model.User{}, errors.New("email not found or invalid in token claims")
	}

	return user, nil
}
