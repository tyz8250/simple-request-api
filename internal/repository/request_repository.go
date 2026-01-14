package repository

import "github.com/ricosh/simple-request-api/internal/model"

// RequestRepository は、リクエストの保存方法と取得方法を定義します。
type RequestRepository interface {
	FindAll() ([]model.Request, error)
	FindByID(id int64) (model.Request, error)
}
