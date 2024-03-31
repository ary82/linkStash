package api

import (
	"encoding/json"
	"log"
	"net/http"

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
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) Run() error {
	router := http.NewServeMux()
	serverConfig := &http.Server{
		Addr:    s.Addr,
		Handler: middleware.Logger(router),
	}

	router.HandleFunc("/", s.NotFound)
	router.HandleFunc("GET /stash", s.getStashHandler)

	log.Println("Startin API on", s.Addr)
	err := serverConfig.ListenAndServe()
	return err
}

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	WriteJsonResponse(w, http.StatusNotFound, map[string]string{"error": "not a valid api path"})
}

func (s *Server) getStashHandler(w http.ResponseWriter, r *http.Request) {
	stashes, err := s.Database.GetPublicStashes()
	if err != nil {
		WriteJsonResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJsonResponse(w, http.StatusOK, stashes)
}
