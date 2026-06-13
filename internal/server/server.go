package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nnaxim/train-radar/internal/handler"
)

type Server struct {
	router *chi.Mux
}

func New(stationHandler *handler.StationHandler) *Server {
	router := chi.NewRouter()

	router.Get(
		"/health",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write(([]byte("OK")))
		},
	)

	router.Get(
		"/stations/{id}",
		stationHandler.GetByID,
	)

	return &Server{
		router: router,
	}
}

func (s *Server) Router() http.Handler {
	return s.router
}
