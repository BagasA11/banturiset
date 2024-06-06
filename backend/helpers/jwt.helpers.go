package helpers

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/bagasa11/banturiset/config"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

const penyunting string = "penyunting"
const peneliti string = "peneliti"
const donatur string = "donatur"

func GenerateToken(userID uint, email string, role string, roleID uint) (string, error) {

	r := ""
	// penyunting
	if slices.Contains([]string{"penyunting", "admin", "penelaah", "reviewer"}, strings.ToLower(role)) {
		r = penyunting
	}
	// peneliti
	if slices.Contains([]string{"peneliti", "ilmuwan", "researcher"}, strings.ToLower(role)) {
		r = peneliti
	}
	if slices.Contains([]string{"donatur", "investor", "penyumbang"}, strings.ToLower(role)) {
		r = donatur
	}

	fmt.Println("role: ", r)
	claims := &config.JwtClaims{
		ID:     userID,
		Email:  email,
		Role:   r,
		RoleID: roleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(config.JWT_KEY)
	if err != nil {
		fmt.Println("error when create jwt.. ", err)
		log.Fatal(err)
		return "", err
	}
	return accessToken, err
}

func ValidateToken(inputToken string) (*config.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(inputToken, &config.JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_KEY), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*config.JwtClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func UpdateToken(oldToken string) (string, error) {
	claims, err := ValidateToken(oldToken)
	if err != nil {
		return "", err
	}

	newToken, err := GenerateToken(claims.ID, claims.Email, claims.Role, claims.RoleID)
	if err != nil {
		return "", err
	}
	return newToken, nil
}
