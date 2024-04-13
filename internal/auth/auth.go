package auth

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/api/idtoken"
)

type ContextKey string

type ContextVal struct {
	UserId int
	Email  string
}

type CustomClaims struct {
	UserId int `json:"uid"`
	jwt.RegisteredClaims
}

func GenerateJWT(UserId int, email string) (string, error) {
	claims := &CustomClaims{
		UserId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    os.Getenv("ISSUER"),
			Subject:   email,
		},
	}

	key := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	return ss, err
}

// Function for getting payload from google identity's jwt
func GetData(tokenStr []byte) (*idtoken.Payload, error) {
	payload, err := idtoken.Validate(context.Background(), string(tokenStr), os.Getenv("AUDIENCE"))
	if err != nil {
		return nil, err
	}
	return payload, nil
}
