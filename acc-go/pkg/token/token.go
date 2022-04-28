package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	userId string `json:"userId"`
	jwt.StandardClaims
}

const jwtSecret = "232323"

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()                    //当前时间
	expireTime := nowTime.Add(3 * time.Hour) //有效时间

	claims := Claims{
		"2222",
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "its me",
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
