package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ary82/urlStash/database"
	"github.com/ary82/urlStash/middleware"
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
		Handler: middleware.Logger(router),
	}

	router.HandleFunc("/", s.NotFound)
	router.HandleFunc("GET /stash", s.getStashHandler)
	router.HandleFunc("POST /login", s.login)

	log.Println("Startin API on", s.Addr)
	err := serverConfig.ListenAndServe()
	return err
}

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	WriteJsonResponse(w, http.StatusNotFound, map[string]string{"error": "not a valid api path"})
}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	tokenStr, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJsonErr(w, err)
		return
	}
	defer r.Body.Close()

	// Get the token Payload from Google idToken JWT
	payload, err := middleware.GetData(tokenStr)
	if err != nil {
		WriteJsonErr(w, err)
		return
	}

	// Insert or Update User from Google's information
	err = s.Database.InsertUserByPayload(payload)
	if err != nil {
		WriteJsonErr(w, err)
		return
	}

	// Get User by their email
	user, err := s.Database.GetUserByEmail(payload.Claims["email"].(string))
	if err != nil {
		WriteJsonErr(w, err)
		return
	}

	WriteJsonResponse(w, http.StatusOK, user)
}

func (s *Server) getStashHandler(w http.ResponseWriter, r *http.Request) {
	stashes, err := s.Database.GetPublicStashes()
	if err != nil {
		WriteJsonResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJsonResponse(w, http.StatusOK, stashes)
}
