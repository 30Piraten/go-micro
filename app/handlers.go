package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Creating Handlers
// Defining the response and request struct
type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func main() {
	port := 7575

	// Creating the validation handler with the HelloWorldHandler
	handler := newValidationHandler(newHelloWorldHandler())

	http.Handle("/world", handler)

	log.Printf("Serving trekking on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// Validation handler definition
type validationHandler struct {
	next http.Handler
}

// Constructor for validationHandler
func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

// ServeHTTP method for validationHandler
func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest

	// Check that the Content-Type is application/json
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Content-Type should be 'application/json'", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)

	// Decode the JSON body into the request struct
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Pass the request to the next handler
	h.next.ServeHTTP(rw, r)
}

// HelloWorldHandler definition
type helloWorldHandler struct{}

// Constructor for helloWorldHandler
func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

// ServeHTTP method for helloWorldHandler
func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// rw.Header().Set("Content-Type", "application/json")

	response := helloWorldResponse{Message: "Di dove sei?"}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

// Run server => go run handlers.go first
// curl -X POST -H "Content-Type: application/json" -d '{"name": "your_name"}' http://localhost:7575/world
