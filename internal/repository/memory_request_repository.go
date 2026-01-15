package repository

import (
	"errors"

	"github.com/ricosh/simple-request-api/internal/model"
)

// MemoryRequestRepository は RequestRepository のインメモリ実装です。
type MemoryRequestRepository struct {
	requests []model.Request
}

// NewMemoryRequestRepository は新しい MemoryRequestRepository を作成します。
func NewMemoryRequestRepository() *MemoryRequestRepository {
	return &MemoryRequestRepository{
		requests: []model.Request{
			{
				ID:          1,
				Title:       "住民票申請",
				Description: "引っ越しに伴う申請",
				Status:      "pending",
			},
		},
	}
}

// FindAll returns all requests.
func (r *MemoryRequestRepository) FindAll() ([]model.Request, error) {
	return r.requests, nil
}

// FindByID returns a request by ID.
func (r *MemoryRequestRepository) FindByID(id int64) (model.Request, error) {
	for _, req := range r.requests {
		if req.ID == id {
			return req, nil
		}
	}
	return model.Request{}, errors.New("request not found")
}
