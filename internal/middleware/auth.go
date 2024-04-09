package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/api/idtoken"
)

type ContextKey string

func Auth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the cookie
		cookie, err := r.Cookie("urlstashJwt")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			// Validate the algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Get subject from claims
			subject, err := claims.GetSubject()
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			// Put the claims in the request's context
			// Can get by r.Context().Value(ctxKey).(jwt.MapClaims) if putting claim
			// Or simply by r.Context().Value(ctxKey) if simple value like string
			ctxKey := ContextKey("username")
			ctx := context.WithValue(r.Context(), ctxKey, subject)

			// Serve the next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
	}
}

func GenerateJWT(username string) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    os.Getenv("ISSUER"),
		Subject:   username,
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
