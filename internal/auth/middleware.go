package auth

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ary82/urlStash/internal/database"
	"github.com/ary82/urlStash/internal/utils"
)

// Authn middleware
func AuthMiddleware(optional bool, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the cookie
		cookie, err := r.Cookie("urlstashJwt")
		if err != nil {
			utils.ClearJwtCookie(w, "urlstashJwt")
			// If Auth is optional for this route, serve next
			if optional {
				next.ServeHTTP(w, r)
			} else {
				utils.WriteJsonUnauthorized(w, err)
			}
			return
		}

		// Get a struct with email, userid from jwt
		contextVal, err := ParseJWT(cookie.Value)
		if err != nil {
			utils.ClearJwtCookie(w, "urlstashJwt")
			// If Auth is optional for this route, serve next
			if optional {
				next.ServeHTTP(w, r)
			} else {
				utils.WriteJsonUnauthorized(w, err)
			}
			return
		}

		// Put the claims in the request's context
		// Can get by r.Context().Value(ctxKey).(TYPE ASSERTION)
		ctxKey := ContextKey("user")
		ctx := context.WithValue(r.Context(), ctxKey, contextVal)

		// Serve the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// Authz middleware that allows access either if user owns this stash,
// or the stash is public. Takes in stash from the url path.
func AuthzStash(
	database database.DB,
	next http.Handler,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		pathStr := r.PathValue("id")
		stashId, err := strconv.Atoi(pathStr)
		if err != nil {
			utils.WriteJsonResponse(
				w,
				http.StatusBadRequest,
				map[string]string{"error": "can't convert path to int"},
			)
			return
		}

		// Get current user from context
		currentUser, ok := r.Context().Value(ContextKey("user")).(*ContextVal)

		// check if currentUser is owner
		var isOwner bool
		if currentUser != nil && ok {
			owner, err := database.CheckOwner(currentUser.UserId, stashId)
			isOwner = owner
			if err != nil {
				utils.WriteJsonServerErr(w, err)
				return
			}
		}

		// check if stash is public
		isPublic, err := database.CheckStashPublic(stashId)
		if err != nil {
			utils.WriteJsonServerErr(w, err)
			return
		}

		if isOwner || isPublic {
			// Serve the next handler
			next.ServeHTTP(w, r)
		} else {
			utils.WriteJsonUnauthorized(w, fmt.Errorf("Not Allowed"))
			return
		}
	}
}
