package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nnaxim/train-radar/internal/service"
)

type StationHandler struct {
	service *service.StationService
}

func NewStationHandler(
	service *service.StationService,
) *StationHandler {
	return &StationHandler{
		service: service,
	}
}

func (h *StationHandler) GetByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	idParam := chi.URLParam(
		r,
		"id",
	)

	id, err := strconv.ParseUint(
		idParam,
		10,
		64,
	)

	if err != nil {
		http.Error(
			w,
			"invalid station id",
			http.StatusBadRequest,
		)

		return
	}

	station, err := h.service.GetByID(
		r.Context(),
		id,
	)

	if err != nil {
		http.Error(
			w,
			"station not found",
			http.StatusNotFound,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(station)
}
