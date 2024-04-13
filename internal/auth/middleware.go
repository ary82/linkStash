package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/ary82/urlStash/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the cookie
		cookie, err := r.Cookie("urlstashJwt")
		if err != nil {
			ClearJwtCookie(w)
			utils.WriteJsonUnauthorized(w, err)
			return
		}

		token, err := jwt.ParseWithClaims(
			cookie.Value,
			&CustomClaims{},
			func(token *jwt.Token) (interface{}, error) {
				// Validate the algorithm
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {

			// Get email from claims
			email, err := claims.GetSubject()
			if err != nil {
				utils.WriteJsonServerErr(w, err)
				return
			}

			// User struct for use in context
			contextVal := &ContextVal{
				UserId: claims.UserId,
				Email:  email,
			}

			// Put the claims in the request's context
			// Can get by r.Context().Value(ctxKey).(TYPE ASSERTION)
			ctxKey := ContextKey("user")
			ctx := context.WithValue(r.Context(), ctxKey, contextVal)

			// Serve the next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			ClearJwtCookie(w)
			utils.WriteJsonUnauthorized(w, err)
		}
	}
}
