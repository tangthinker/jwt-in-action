package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	ExpirationTime = 60
	Secret         = "ijfow34f49129"
) // 60 minutes

func GenerateJWT(userId string) (string, error) {

	claim := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(ExpirationTime)).Unix(),
		Id:        userId,
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func VerifyJWT(tokenString string) (string, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if err != nil {

		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Id, nil

}
