package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mohammadgh1370/url-shortner/internal/config"
	"time"
)

type JWTClaim struct {
	Identifier string `json:"identifier"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(identifier string) (tokenString string, err error) {
	expirationTime := time.Now().Add(time.Duration(config.JWT_ACCESS_EXP_TIME) * time.Minute)
	claims := &JWTClaim{
		Identifier: identifier,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(config.JWT_SECRET_KEY))
	return
}

func GenerateRefreshToken(identifier string) (tokenString string, err error) {
	expirationTime := time.Now().Add(time.Duration(config.JWT_REFRESH_EXP_TIME) * time.Minute)
	claims := &JWTClaim{
		Identifier: identifier,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(config.JWT_SECRET_KEY))
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWT_SECRET_KEY), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
