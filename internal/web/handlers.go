package web

import (
	"encoding/json"
	"gokyoapi/internal/service"
	"net/http"
	"strconv"
)

type GokyoHandlers struct {
	service *service.GokyoService
}

func NewGokyoHandler(service *service.GokyoService) *GokyoHandlers {
	return &GokyoHandlers{service: service}
}

func (h *GokyoHandlers) GetStatus(w http.ResponseWriter, r *http.Request) {
	IsUp, err := h.service.GetStatus()
	if err != nil {
		http.Error(w, "Failed to get status", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(IsUp)
}

func (h *GokyoHandlers) GetGokyo(w http.ResponseWriter, r *http.Request) {
	gokyo, err := h.service.GetGokyo()
	if err != nil {
		http.Error(w, "Failed to get Gokyo", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gokyo)
}

func (h *GokyoHandlers) CreateGokyo(w http.ResponseWriter, r *http.Request) {
	var gokyo service.Gokyo
	err := json.NewDecoder(r.Body).Decode(&gokyo)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	err = h.service.CreateGokyo(&gokyo)
	if err != nil {
		http.Error(w, "invalid request", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(gokyo)
}

func (h *GokyoHandlers) DeleteGokyo(w http.ResponseWriter, r *http.Request) {
	Str := r.PathValue("id")
	id, err := strconv.Atoi(Str)
	if err != nil {
		http.Error(w, "Invalid GokyoId ", http.StatusBadRequest)
	}
	if err := h.service.DeleteGokyo(id); err != nil {
		http.Error(w, "failed to delete book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *GokyoHandlers) UpdateGokyo(w http.ResponseWriter, r *http.Request) {
	var gokyo service.Gokyo
	err := json.NewDecoder(r.Body).Decode(&gokyo)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	err = h.service.UpdateGokyo(&gokyo)
	if err != nil {
		http.Error(w, "invalid request", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(gokyo)
}
