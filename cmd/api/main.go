package main

import (
	"fmt"
	"time"

	"github.com/ricosh/simple-request-api/internal/model"
)

func main() {
	req := model.Request{
		ID:          1,
		Title:       "住民票の申請",
		Description: "引っ越しに伴う申請",
		Status:      "pending",
		CreatedAt:   time.Now(),
	}
	fmt.Println(req)
}
