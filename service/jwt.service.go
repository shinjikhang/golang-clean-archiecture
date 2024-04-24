package service

import (
	"clean-architecture/util"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JwtService interface {
	GenerateTokenPair(userId interface{}) (map[string]string, error)
}

type jwtCustomClaim struct {
	UserId interface{} `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	issuer    string
	secretKey string
}

func NewJwtService() *jwtService {
	return &jwtService{
		issuer: os.Getenv("TOKEN_ISSUER"), secretKey: util.GetSecretKey(),
	}
}

func (service *jwtService) getTokenClaims(
	userId interface{}, expiryDays int,
) *jwtCustomClaim {
	return &jwtCustomClaim{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, expiryDays).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
}

func (service *jwtService) GenerateTokenPair(userId interface{}) (map[string]string, error) {
	tokenClaims := service.getTokenClaims(userId, 15)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		return nil, err
	}
	refreshTokenClaims := service.getTokenClaims(userId, 30)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(service.secretKey))
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"access_token":  tokenString,
		"refresh_token": refreshTokenString,
	}, nil
}
