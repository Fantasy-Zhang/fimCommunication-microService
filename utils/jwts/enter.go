package jwts

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtPayload struct {
	UserId   uint   `json:"user_id"`
	Nickname string `json:"nickname"` //用户名
	Role     int    `json:"role"`     //权限 1:管理员 2:普通用户
}
type CustomClaims struct {
	JwtPayload
	jwt.RegisteredClaims //
}

func GenToken(payload JwtPayload, accessSecret string, expires int) (string, error) {
	claims := CustomClaims{
		payload,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expires))),
			//IssuedAt:  jwt.NewNumericDate(NowTime()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessSecret))
}

// 解析token
func ParseToken(tokenString string, accessSecret string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
