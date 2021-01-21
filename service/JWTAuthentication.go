package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 认证接口
type JWTService interface {
	CreateToken(username string) string
	ValidateToken(token string) (*jwt.Token, error)
}

// 认证结构体
type jwtService struct {
	secret  string
	issurer string
}

type JwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

// 实例化认证结构体
func NewJWAuthService() JWTService {
	return &jwtService{
		secret:  "secret", //自定义配置
		issurer: "lsrong",
	}
}

// 创建认证token
func (service *jwtService) CreateToken(username string) string {
	claims := &JwtCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issurer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(service.secret))
	if err != nil {
		panic(fmt.Errorf("Create signed tokon failure error:%v ", err))
	}
	return token
}

// 校验token
func (service *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v ", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(service.secret), nil
	})
}
