package handler

import (
	"encoding/json"
	"net/http"

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
