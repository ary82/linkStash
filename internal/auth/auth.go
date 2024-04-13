package auth

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/api/idtoken"
)

// Custom jwt claims for urlstash
type CustomClaims struct {
	UserId int `json:"uid"`
	jwt.RegisteredClaims
}

// Useful data from Google's JWT
type GoogleData struct {
	Username string
	Name     string
	Email    string
	Picture  string
}

// Context key/val pair
type ContextKey string
type ContextVal struct {
	UserId int
	Email  string
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

// Function for extracting payload from google identity's jwt
func GetPayload(tokenStr []byte) (*idtoken.Payload, error) {
	payload, err := idtoken.Validate(
		context.Background(),
		string(tokenStr),
		os.Getenv("AUDIENCE"),
	)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

// Extract data from google jwt's payload
func GetData(p *idtoken.Payload) *GoogleData {
	return &GoogleData{
		// Generate username from email for now
		// TODO: Let users pick
		Username: strings.Split(p.Claims["email"].(string), "@")[0],
		Email:    p.Claims["email"].(string),
		Name:     p.Claims["name"].(string),
		Picture:  p.Claims["picture"].(string),
	}
}

// Clears the login jwt cookiw from client
func ClearJwtCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     "urlstashJwt",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
}
