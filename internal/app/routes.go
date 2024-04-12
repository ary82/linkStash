package app

import (
	"net/http"

	"github.com/ary82/urlStash/internal/auth"
)

func (s *Server) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("/", s.NotFound)
	router.HandleFunc("GET /stash", s.getPublicStashesHandler)
	router.HandleFunc("GET /stashes/{user}", s.getUserStashesHandler)
	router.HandleFunc("GET /stash/{id}", s.getStashHandler)
	router.HandleFunc("POST /login", s.LoginHandler)

  router.HandleFunc("GET /user/{id}", s.getUserHandler)

	// These routes require authenticaltion
	router.Handle("POST /logout", auth.AuthMiddleware(http.HandlerFunc(s.LogoutHandler)))
	router.Handle("GET /private", auth.AuthMiddleware(http.HandlerFunc(s.getPrivate)))

}
