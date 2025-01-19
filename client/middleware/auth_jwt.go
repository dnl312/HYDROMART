package helpers

import (
	"client/model"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetUserIDFromToken(c echo.Context) (model.User, error) {
	userToken := c.Get("user")
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
		return model.User{}, errors.New("User ID not found or invalid in token claims")
	}
	_, ok = claims["role"].(string)
	if !ok {
		return model.User{}, errors.New("Role not found or invalid in token claims")
	}
	_, ok = claims["email"].(string)
	if !ok {
		return model.User{}, errors.New("Email not found or invalid in token claims")
	}

	return user, nil
}
