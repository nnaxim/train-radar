package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nnaxim/train-radar/internal/dto"
	"github.com/nnaxim/train-radar/internal/model"
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

func (h *StationHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req dto.CreateStationRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(
			w,
			"Invalid request body",
			http.StatusBadRequest,
		)

		return
	}

	station := &model.Station{
		Name:      req.Name,
		Country:   req.Country,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	err = h.service.Create(
		r.Context(),
		station,
	)

	if err != nil {
		http.Error(
			w,
			"failed to create station",
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(
		http.StatusCreated,
	)

	json.NewEncoder(w).Encode(
		station,
	)

}

func (h *StationHandler) GetAll(
	w http.ResponseWriter,
	r *http.Request,
) {
	stations, err := h.service.GetAll(
		r.Context(),
	)

	if err != nil {
		http.Error(
			w,
			"failed to get stations",
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(stations)
}

func (h *StationHandler) DeleteByID(
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

	err = h.service.DeleteByID(
		r.Context(),
		id,
	)

	if err != nil {
		http.Error(
			w,
			"failed to delete station",
			http.StatusInternalServerError,
		)

		return
	}

	w.WriteHeader(
		http.StatusNoContent,
	)
}

func (h *StationHandler) RegisterRoutes(
	r chi.Router,
) {

	r.Get(
		"/stations",
		h.GetAll,
	)
	r.Get(
		"/stations/{id}",
		h.GetByID,
	)

	r.Post(
		"/stations",
		h.Create,
	)

	r.Delete(
		"/stations/{id}",
		h.DeleteByID,
	)

}
