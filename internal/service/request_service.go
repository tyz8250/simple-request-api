package request_service

import (
	"github.com/ricosh/simple-request-api/internal/model"
	"github.com/ricosh/simple-request-api/internal/repository"
)

// RequestService はリクエストに関するビジネスロジックを提供します。
type RequestService struct {
	repo repository.RequestRepository
}

// NewRequestService は、新しい RequestService を作成します。
func NewRequestService(repo repository.RequestRepository) *RequestService {
	return &RequestService{repo: repo}
}

// GetAll は、すべてのリクエストを取得します。
func (s *RequestService) GetAll() ([]model.Request, error) {
	return s.repo.FindAll()
}
