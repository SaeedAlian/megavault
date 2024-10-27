package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/test", h.handleTest).Methods("GET")
}

func (h *Handler) handleTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(200)
	data := map[string]string{"message": "hello world"}

	json.NewEncoder(w).Encode(data)
}
