package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ricosh/simple-request-api/internal/model"
)

// RequestHandler は、Requestに関するHTTPリクエストを処理する。
type RequestHandler struct {
}

// GetRequests は、Requestのリストを取得する。
func (h *RequestHandler) GetRequests(w http.ResponseWriter, r *http.Request) {
	requests := []model.Request{
		{
			ID:          1,
			Title:       "住民票の申請",
			Description: "引っ越しに伴う申請",
			Status:      "pending",
			CreatedAt:   time.Now(),
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(requests)
}
