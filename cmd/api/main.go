package main

import (
	"log"
	"net/http"

	"github.com/ricosh/simple-request-api/internal/handler"
	request_service "github.com/ricosh/simple-request-api/internal/service"
)

func main() {
	requestService := &request_service.RequestService{}
	requestHandler := handler.NewRequestHandler(requestService)

	http.HandleFunc("/requests", requestHandler.GetRequests)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
