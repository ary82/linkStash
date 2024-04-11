package app

import (
	"net/http"

	"github.com/ary82/urlStash/internal/middleware"
)

func (s *Server) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("/", s.NotFound)
	router.HandleFunc("GET /stash", s.getPublicStashHandler)
	router.HandleFunc("GET /stash/{id}", s.getStashHandler)
	router.HandleFunc("POST /login", s.LoginHandler)

	// These routes require authenticaltion
	router.Handle("POST /logout", middleware.Auth(http.HandlerFunc(s.LogoutHandler)))
	router.Handle("GET /private", middleware.Auth(http.HandlerFunc(s.getPrivate)))

}
