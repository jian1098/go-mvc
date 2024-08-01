package common

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cast"
)

// Payload 载荷
type JwtClaims struct {
	jwt.StandardClaims
	UserId int
}

// JWT数据
type Jwt struct {
	ExpiresAt int64
	Token     string
}

// 生成JWT
func CreateJwt(userId int) (*Jwt, error) {
	mySigningKey := []byte(os.Getenv("JWT_SECRET"))

	// Create the Claims
	expiresAt := time.Now().Add(time.Second * time.Duration(cast.ToInt(os.Getenv("JWT_EXPIRE_TIME")))).Unix() // 过期时间（秒）
	claims := &JwtClaims{}
	claims.ExpiresAt = expiresAt
	claims.UserId = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return nil, err
	}

	//返回数据
	jwtData := &Jwt{}
	jwtData.Token = tokenString
	jwtData.ExpiresAt = expiresAt

	return jwtData, err
}

// 解析JWT
func ParseJwt(tokenString string) (*JwtClaims, bool) {
	//解析令牌字符串
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Println(err)
		return nil, false
	}

	//获取载荷
	claims, ok := token.Claims.(*JwtClaims)
	if ok && token.Valid {
		return claims, true
	}

	return nil, false
}
