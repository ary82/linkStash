package app

import (
	"net/http"

	"github.com/ary82/urlStash/internal/auth"
)

func (s *Server) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("/", s.NotFound)
	router.HandleFunc("GET /stash", s.getPublicStashesHandler)
	router.HandleFunc("GET /stash/{id}", s.getStashHandler)
	router.HandleFunc("GET /user/{id}", s.getUserProfileHandler)
	router.HandleFunc("POST /login", s.LoginHandler)
	router.Handle("POST /logout", http.HandlerFunc(s.LogoutHandler))

	// These routes require authenticaltion
	router.Handle("GET /private", auth.AuthMiddleware(http.HandlerFunc(s.getPrivate)))
	router.Handle("GET /me", auth.AuthMiddleware(http.HandlerFunc(s.getUserHandler)))

}
