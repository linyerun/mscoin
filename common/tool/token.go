package tool

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"
)

var TokenInValidError = errors.New("token已过期了")

func GenerateToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func ParseToken(tokenString string, secret string) (userId int64, err error) {
	// 解析Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if err != nil {
		return -1, err
	}

	// 获取token信息(TODO: 最终为啥int64被解析成了float64了)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp := int64(claims["exp"].(float64)); exp <= time.Now().Unix() {
			return -1, TokenInValidError
		}

		return int64(claims["userId"].(float64)), nil
	}

	return -1, TokenInValidError
}
