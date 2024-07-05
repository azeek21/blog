package service

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"time"

	"github.com/azeek21/blog/models"
	"github.com/golang-jwt/jwt/v5"
)

type JwtServ struct{}

func NewJwtSerice() JwtService {
	return JwtServ{}
}

func (s JwtServ) CreateJwt(user *models.User) (string, error) {
	token := ""
	key, err := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	if err != nil {
		return token, err

	}
	generatedAt := time.Now().Unix()
	expiresAt := generatedAt + int64((time.Hour * 720).Seconds())
	t := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"id":    user.ID,
		"iss":   "blog.askaraliev.uz",
		"iat":   generatedAt,
		"exp":   expiresAt,
		"email": user.Email,
	})
	token, err = t.SignedString(key)
	if err != nil {
		return token, err

	}
	return token, nil

}

func (s JwtServ) VerifyJwt(token string) (uint, error) {
	return 0, nil
}
