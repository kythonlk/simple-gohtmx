package api

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwt auth token gen and validation

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(username string, userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
