package helpers

/*
Feature 1: validate token

header checker

validate token err cases:
cases 1: expired token
  alert:"token expired"
  do: update token
cases 2: invalid token and others

check token in memcache:
case 1: exist
case 2: not exist

return claims
--------------------------------
feature 2: update token
*/

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"

	"time"

	"github.com/bagasa11/banturiset/config"

	"github.com/golang-jwt/jwt/v5"
)

// var tokenCache = config.GetCacheTTL()

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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(4 * time.Hour)),
		},
	}

	// hashing token
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(config.JWT_KEY)
	if err != nil {
		fmt.Println("error when create jwt.. ", err)
		log.Fatal(err)
		return "", err
	}

	// insert token to cache
	if err = insertCacheToken(claims.ID, accessToken); err != nil {
		fmt.Println("failed insert token to cache")
		return "", errors.New("failed insert token to cache")
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

	claims, ok := token.Claims.(*config.JwtClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	cchTkn, err := getCachedToken(claims.ID)
	if err != nil {
		if err == config.CacheNotFound {
			return nil, errors.New("unauthorized token: token hasn't been registered in session")
		}
		return nil, err
	}
	if cchTkn != inputToken {
		return nil, errors.New("session token have been tampered or outdated")
	}
	return claims, nil
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

	if err = updateCachedToken(claims.ID, newToken); err != nil {
		return "", err
	}

	return newToken, nil
}

func insertCacheToken(id uint, token string) error {
	c := config.GetCacheTTL()
	var key = fmt.Sprintf("user-%d", id)
	return c.SetWithTTL(key, token, time.Duration((4*time.Hour)+(10*time.Minute)))
}

func getCachedToken(id uint) (token string, err error) {
	c := config.GetCacheTTL()
	key := fmt.Sprintf("user-%d", id)
	t, err := c.Get(key)
	if err != nil {
		if err == config.CacheNotFound {
			return "", config.CacheNotFound
		}
		return "", err
	}
	return t.(string), nil
}

func rmCachedToken(id uint) error {
	c := config.GetCacheTTL()
	key := fmt.Sprintf("user-%d", id)
	return c.Remove(key)
}

func updateCachedToken(id uint, newToken string) error {

	// check if token exist, then remove that token
	if err := rmCachedToken(id); err == config.CacheNotFound {
		return errors.New("unauthorized token: token hasn't been registered in session")
	}
	// insert token to memcache
	if err := insertCacheToken(id, newToken); err != nil {
		return err
	}
	return nil
}
