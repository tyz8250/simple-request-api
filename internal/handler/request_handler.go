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
