package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	request_service "github.com/ricosh/simple-request-api/internal/service"
)

// RequestHandler は、Requestに関するHTTPリクエストを処理する。
type RequestHandler struct {
	service *request_service.RequestService
}

type createRequestInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// NewRequestHandler は、RequestHandlerを生成する。
func NewRequestHandler(s *request_service.RequestService) *RequestHandler {
	return &RequestHandler{service: s}
}

// GetRequests は、Requestのリストを取得する。
func (h *RequestHandler) GetRequests(w http.ResponseWriter, r *http.Request) {
	requests, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "failed to get requests", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(requests)
}

// CreateRequest は、新しいRequestを保存する。
func (h *RequestHandler) CreateRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input createRequestInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	req, err := h.service.Create(input.Title, input.Description)
	if err != nil {
		http.Error(w, "failed to create request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}
func (h *RequestHandler) GetRequestByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// /requests/{id} から id を取り出す
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "invalid path", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	req, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "request not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(req)
}
