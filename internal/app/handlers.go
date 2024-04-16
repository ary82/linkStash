package app

import (
	"io"
	"net/http"
	"strconv"

	"github.com/ary82/urlStash/internal/auth"
	"github.com/ary82/urlStash/internal/utils"
)

func (s *Server) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteJsonResponse(
		w,
		http.StatusNotFound,
		map[string]string{"error": "not a valid api path"},
	)
}

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	utils.ClearJwtCookie(w, "urlstashJwt")
	utils.WriteJsonResponse(
		w,
		http.StatusOK,
		map[string]string{"message": "logged out"},
	)
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	tokenStr, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteJsonServerErr(w, err)
		return
	}
	defer r.Body.Close()

	jwt, err := auth.Login(tokenStr, s.Database)
	if err != nil {
		utils.WriteJsonServerErr(w, err)
	}

	// Serve the jwt as cookie
	utils.SetCookie(w, "urlstashJwt", *jwt)

	utils.WriteJsonResponse(w, http.StatusOK, map[string]string{
		"message": "Successfully logged in",
	})
}

func (s *Server) getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get email from suth context
	email := r.Context().Value(auth.ContextKey("user")).(*auth.ContextVal).Email

	user, err := s.Database.GetUserByEmail(email)
	if err != nil {
		utils.WriteJsonServerErr(w, err)
		return
	}
	utils.WriteJsonResponse(w, http.StatusOK, user)
}

func (s *Server) getUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	pathStr := r.PathValue("id")
	pathInt, err := strconv.Atoi(pathStr)
	if err != nil {
		utils.WriteJsonBadReq(w, err)
		return
	}
	user, err := s.Database.GetUserProfile(pathInt)
	if err != nil {
		utils.WriteJsonServerErr(w, err)
		return
	}
	utils.WriteJsonResponse(w, http.StatusOK, user)

}

func (s *Server) getPublicStashesHandler(w http.ResponseWriter, r *http.Request) {
	stashes, err := s.Database.GetPublicStashes()
	if err != nil {
		utils.WriteJsonServerErr(w, err)
		return
	}
	utils.WriteJsonResponse(w, http.StatusOK, stashes)
}

func (s *Server) getStashHandler(w http.ResponseWriter, r *http.Request) {
	pathStr := r.PathValue("id")
	pathInt, err := strconv.Atoi(pathStr)
	if err != nil {
		utils.WriteJsonBadReq(w, err)
		return
	}
	stash, err := s.Database.GetStashDetailed(pathInt)
	if err != nil {
		utils.WriteJsonServerErr(w, err)
		return
	}
	utils.WriteJsonResponse(w, http.StatusOK, stash)
}

// Only for testing auth routes
// TODO: delete this
func (s *Server) getPrivate(w http.ResponseWriter, r *http.Request) {
	utils.WriteJsonResponse(w, http.StatusOK, map[string]string{"message": "private path accessed"})
}
