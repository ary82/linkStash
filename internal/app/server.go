package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ary82/urlStash/internal/auth"
	"github.com/ary82/urlStash/internal/database"
	"github.com/ary82/urlStash/internal/logging"
)

type Server struct {
	Addr     string
	Database *database.DB
}

func NewApiServer(addr string, database *database.DB) *Server {
	return &Server{
		Addr:     addr,
		Database: database,
	}
}

func WriteJsonResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Add("Content-Type", "application/json")
	if os.Getenv("MODE") == "DEV" {
		w.Header().Add("Access-Control-Allow-Origin", os.Getenv("CLIENT_URL"))
		w.Header().Add("Access-Control-Allow-Credentials", "true")
	}
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func WriteJsonErr(w http.ResponseWriter, err error) {
	WriteJsonResponse(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
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

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	WriteJsonResponse(w, http.StatusNotFound, map[string]string{"error": "not a valid api path"})
}

func (s *Server) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Context().Value(auth.ContextKey("username")))
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
	WriteJsonResponse(w, http.StatusOK, map[string]string{"message": "logged out"})
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	tokenStr, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJsonErr(w, err)
		return
	}
	defer r.Body.Close()

	// Get the token Payload from Google idToken JWT
	payload, err := auth.GetData(tokenStr)
	if err != nil {
		WriteJsonErr(w, err)
		return
	}
	fmt.Println(payload.Claims)

	// Insert or Update User from Google's information
	err = s.Database.UpsertUserByPayload(payload)
	if err != nil {
		WriteJsonErr(w, err)
		return
	}

	jwt, err := auth.GenerateJWT(strings.Split(payload.Claims["email"].(string), "@")[0])
	if err != nil {
		WriteJsonErr(w, err)
		return
	}

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

	// Also return the picture in response
	WriteJsonResponse(w, http.StatusOK, map[string]string{
		"message": "Successfully logged in",
		"picture": payload.Claims["picture"].(string),
	})
}

func (s *Server) getPublicStashesHandler(w http.ResponseWriter, r *http.Request) {
	stashes, err := s.Database.GetPublicStashes()
	if err != nil {
		WriteJsonErr(w, err)
		return
	}
	WriteJsonResponse(w, http.StatusOK, stashes)
}

func (s *Server) getUserStashesHandler(w http.ResponseWriter, r *http.Request) {
	pathStr := r.PathValue("user")
	stashes, err := s.Database.GetStashesByUser(pathStr)
	if err != nil {
		WriteJsonErr(w, err)
		return
	}
	WriteJsonResponse(w, http.StatusOK, stashes)
}

func (s *Server) getUserHandler(w http.ResponseWriter, r *http.Request) {
	pathStr := r.PathValue("id")
	pathInt, err := strconv.Atoi(pathStr)
	if err != nil {
		WriteJsonResponse(
			w,
			http.StatusBadRequest,
			map[string]string{"error": "can't convert path to int"},
		)
		return
	}
	user, err := s.Database.GetUserProfile(pathInt)
	if err != nil {
		WriteJsonErr(w, err)
		return
	}
	WriteJsonResponse(w, http.StatusOK, user)

}

func (s *Server) getStashHandler(w http.ResponseWriter, r *http.Request) {
	pathStr := r.PathValue("id")
	pathInt, err := strconv.Atoi(pathStr)
	if err != nil {
		WriteJsonResponse(
			w,
			http.StatusBadRequest,
			map[string]string{"error": "can't convert path to int"},
		)
		return
	}
	stash, err := s.Database.GetStash(pathInt)
	if err != nil {
		WriteJsonErr(w, err)
		return
	}
	WriteJsonResponse(w, http.StatusOK, stash)
}

func (s *Server) getPrivate(w http.ResponseWriter, r *http.Request) {
	WriteJsonResponse(w, http.StatusOK, map[string]string{"message": "private path accessed"})
}
