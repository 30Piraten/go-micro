package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type resServe struct {
	Phrase         string `json:"phrase"`
	Identification int    `json:"identification"`
}

func main() {

	port := 2325

	http.DefaultServeMux.HandleFunc("/", home)
	log.Printf("Serving at 127.0.0.1:%v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

	// fileServer route function
	imageHandler := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images/", imageHandler))
}

func home(w http.ResponseWriter, r *http.Request) {

	res := resServe{
		Phrase:         "Confirm the identification number...",
		Identification: 220567,
	}

	w.Header().Set("Content-type", "application/json")

	encoder := json.NewEncoder(w)
	err := encoder.Encode(&res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}
