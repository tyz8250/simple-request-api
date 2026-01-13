package main

import (
	"log"
	"net/http"

	"github.com/ricosh/simple-request-api/internal/handler"
)

func main() {
	requestHandler := &handler.RequestHandler{}

	http.HandleFunc("/requests", requestHandler.GetRequests)
	log.Println("Server started af :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
