package utils

import (
	"fmt"
	"github.com/doutokk/doutok/app/auth/conf"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// 加密密钥
var jwtSecret = []byte(conf.GetConf().JwtSecret)

// 自定义负载
type CustomClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// 生成 jwt
func GenerateJWT(userID int, expirationTime time.Duration) (string, error) {
	// 设置负载
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			// 签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
			// 签发人
			Issuer: "DouTok",
		},
	}

	// 创建 token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 校验 jwt
func ValidateJWT(tokenString string) (*CustomClaims, error) {

	// 加上处理 Bearer 的逻辑，删除 Bearer
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// 解析并校验 token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 校验签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	// 签名方法不对
	if err != nil {
		return nil, err
	}

	// 提取 claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
