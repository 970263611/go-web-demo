package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"project/consts"
	"project/file"
	"time"
)

type Claims struct {
	UUID string
	jwt.StandardClaims
}

func GenerateJwtToken(claims Claims) (string, error) {
	//让token不过期，通过uuid来校验过期
	expireTime := time.Now().Add(12 * time.Hour)
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer:    consts.TokenIssuer,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(file.GetEnvParam().TokenSecret))
	return token, err
}

func ParseJwtToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(file.GetEnvParam().TokenSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
