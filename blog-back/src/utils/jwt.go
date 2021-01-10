package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type UserClaim struct {
	jwt.StandardClaims

	Id int64
	Username string
}

func GenerateToken(id int64, username, secretKey string, expireSeconds int64) (string, error) {
	claims := &UserClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expireSeconds) * time.Second).Unix(),
		},

		Id:             id,
		Username:       username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ParseToken(tokenStr, secretKey string) (*UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signature method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaim); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("user unauthorized")
	}
}

func GenSecretKey(length int) string {
	result := ""
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	chaLen := len(characters)

	for i:=0; i<length; i++ {
		result += string(characters[rand.Intn(chaLen)])
	}

	return result
}
