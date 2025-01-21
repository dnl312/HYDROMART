package model

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Email   string `json:"email"`
	Role    string `json:"role"`
	UserID  string `json:"user_id"`
	Exp     int64  `json:"exp"`
	jwt.RegisteredClaims
}