package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Claims struct {
	Uid int64  `json:"uid"`
	IP  string `json:"ip"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte("ACC@@Jason.com@@20220606")

func GenerateToken(uid int64, ip string, duration time.Duration) (string, error) {
	expireTime := time.Now().Add(duration) //有效时间
	id := uuid.NewV4()
	claims := Claims{
		Uid: uid,
		IP:  ip,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			ID:        id.String(),
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
