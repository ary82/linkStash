package app

import (
	"net/http"

	"github.com/ary82/urlStash/internal/auth"
)

func (s *Server) RegisterRoutes(router *http.ServeMux) {

	// Default 404
	router.HandleFunc("/", s.notFoundHandler)

	// Get all Public stashes
	router.HandleFunc("GET /stash", s.getPublicStashesHandler)

	// Detailed stash with links/comments
	router.HandleFunc("GET /stash/{id}", s.getStashHandler)

	// Detailed User with their public stashes
	router.HandleFunc("GET /user/{id}", s.getUserProfileHandler)

	// Auth Routes
	router.HandleFunc("POST /login", s.loginHandler)
	router.Handle("POST /logout", http.HandlerFunc(s.logoutHandler))

	// These routes require authenticaltion
	// private path for testing
	router.Handle("GET /private", auth.AuthMiddleware(http.HandlerFunc(s.getPrivate)))
	// Get Current User data
	router.Handle("GET /me", auth.AuthMiddleware(http.HandlerFunc(s.getUserHandler)))

}
