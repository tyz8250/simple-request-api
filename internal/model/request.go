package model

import "time"

// Request は、システムにおける application/request を表します。
// Status は、request の状態を表します。例: "pending", "approved", "rejected"

type Request struct {
	ID          int64
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time // request が作成された日時。自動でjsonに変換される。
}
