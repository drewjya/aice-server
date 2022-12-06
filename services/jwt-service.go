package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(userId string, name string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{secretKey: getSecretKey(), issuer: "ywdagas"}
}

func getSecretKey() string {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "thisISMYSEcretKey"
	}
	return secretKey
}

func (serv *jwtService) GenerateToken(userId string, name string) string {
	claims := &jwtCustomClaim{
		userId,
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    serv.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := tokenJWT.SignedString([]byte(serv.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}
func (serv *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(tL *jwt.Token) (interface{}, error) {
		_, ok := tL.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method %v", tL.Header["alg"])
		}
		return []byte(serv.secretKey), nil
	})
}
