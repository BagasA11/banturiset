package config

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte(os.Getenv("JWT_KEY"))

type JwtClaims struct {
	ID     uint   `json:"id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	RoleID uint   `json:"rid"`
	jwt.RegisteredClaims
}
