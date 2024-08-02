package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloHandlerResponse struct {
	Message    string `json:"message"`
	Occupation string `json:"occupation"`
	Id         int    `json:"id"`
}

func main() {
	port := 2323

	http.HandleFunc("/", helloHandler)

	log.Printf("Server starting on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	res := helloHandlerResponse{
		Message:    "Willkommen",
		Occupation: "Cloud Engineer",
		Id:         203511,
	}

	// set the response header to indicate JSON content
	w.Header().Set("Content-Type", "application/json")

	// encode the response as JSON and
	// write it to the response writer
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
