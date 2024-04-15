package app

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ary82/urlStash/internal/auth"
	"github.com/ary82/urlStash/internal/database"
	"github.com/ary82/urlStash/internal/logging"
	"github.com/ary82/urlStash/internal/utils"
)

type Server struct {
	Addr     string
	Database database.DB
}

func NewApiServer(addr string, database database.DB) *Server {
	return &Server{
		Addr:     addr,
		Database: database,
	}
}

func (s *Server) Run() error {
	router := http.NewServeMux()
	serverConfig := &http.Server{
		Addr:    s.Addr,
		Handler: logging.LoggerMiddleware(router),
	}

	s.RegisterRoutes(router)

	log.Println("Startin API on", s.Addr)
	err := serverConfig.ListenAndServe()
	return err
}

func (s *Server) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteJsonResponse(
		w,
		http.StatusNotFound,
		map[string]string{"error": "not a valid api path"},
	)
}

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	auth.ClearJwtCookie(w)
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

	// Get the token Payload from Google idToken JWT
	payload, err := auth.GetPayload(tokenStr)
	if err != nil {
		utils.WriteJsonServerErr(w, err)
		return
	}

	// Extract useful data form payload
	data := auth.GetData(payload)

	// Insert or Update User from Google's information
	err = s.Database.UpsertUser(
		data.Username,
		data.Name,
		data.Email,
		data.Picture,
	)
	if err != nil {
		utils.WriteJsonServerErr(w, err)
		return
	}

	// Get the upserted User
	user, err := s.Database.GetUserByEmail(data.Email)
	if err != nil {
		utils.WriteJsonServerErr(w, err)
		return
	}

	// Generate the urlStash jwt with claims set as userId and email
	jwt, err := auth.GenerateJWT(user.ID, data.Email)
	if err != nil {
		utils.WriteJsonServerErr(w, err)
		return
	}

	// Serve the jwt as cookie
	cookie := &http.Cookie{
		Name:     "urlstashJwt",
		Value:    jwt,
		Path:     "/",
		MaxAge:   24 * 3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)

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
		utils.WriteJsonResponse(
			w,
			http.StatusBadRequest,
			map[string]string{"error": "can't convert path to int"},
		)
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
		utils.WriteJsonResponse(
			w,
			http.StatusBadRequest,
			map[string]string{"error": "can't convert path to int"},
		)
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
