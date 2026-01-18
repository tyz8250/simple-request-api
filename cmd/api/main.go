package main

import (
	"log"
	"net/http"

	"github.com/ricosh/simple-request-api/internal/handler"
	"github.com/ricosh/simple-request-api/internal/repository"
	request_service "github.com/ricosh/simple-request-api/internal/service"
)

func main() {
	repo := repository.NewMemoryRequestRepository()
	service := request_service.NewRequestService(repo)
	handler := handler.NewRequestHandler(service)

	http.HandleFunc("/requests", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetRequests(w, r)
		case http.MethodPost:
			handler.CreateRequest(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/requests/", handler.GetRequestByID)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
