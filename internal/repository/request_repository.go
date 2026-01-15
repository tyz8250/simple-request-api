package repository

import "github.com/ricosh/simple-request-api/internal/model"

// RequestRepository は、リクエストの保存方法と取得方法を定義。
// interface なので、実装は別ファイルに定義する。(DB実装に依存しない、テストや差し替えを容易にする)

type RequestRepository interface {
	FindAll() ([]model.Request, error)
	FindByID(id int64) (model.Request, error)
}
