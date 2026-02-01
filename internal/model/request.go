package model

import "time"

// Request は、システムにおける application/request を表します。
// Status は、request の状態を表します。例: "pending", "approved", "rejected"

type Request struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
