package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("tHffy9xoTdcCjpqmbk663PqPts7XaLZq")

type Claims struct {
	ID       int64  `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(id int64, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(365 * time.Hour * 24)
	claims := Claims{
		id, username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "GoNav",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
