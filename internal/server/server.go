package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *chi.Mux
}

func New() *Server {
	router := chi.NewRouter()

	router.Get(
		"/health",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write(([]byte("OK")))
		},
	)

	return &Server{
		router: router,
	}
}

func (s *Server) Router() *chi.Mux {
	return s.router
}

func (s *Server) Handler() http.Handler {
	return s.router
}
