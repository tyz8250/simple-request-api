package request_service

import (
	"time"

	"github.com/ricosh/simple-request-api/internal/model"
)

type RequestService struct {
}

// GetAll 関数がすべてのリクエストを取得する。
func (s *RequestService) GetAll() ([]model.Request, error) {
	requests := []model.Request{
		{
			ID:          1,
			Title:       "住民票申請",
			Description: "引っ越しに伴う申請",
			Status:      "pending",
			CreatedAt:   time.Now(),
		},
	}
	return requests, nil
}
