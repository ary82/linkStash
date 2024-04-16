package app

import (
	"log"
	"net/http"

	"github.com/ary82/urlStash/internal/database"
	"github.com/ary82/urlStash/internal/logging"
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
