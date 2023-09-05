package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Question struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Private     bool   `json:"is_private"`
	CreatedBy   string `json:"created_by"`
}

func main() {

	// defining routes
	http.HandleFunc("/ping", pingPong)
	http.HandleFunc("/questions", getQuestions)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:4040", nil))
}

func pingPong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong")
}

func getQuestions(w http.ResponseWriter, r *http.Request) {
	questions := []Question{
		{"Tallest building in South Asia", "1wfags-qsbe-asfg-3d4t", false, "Raj"},
		{"Tallest building in South Asia", "1wfags-qsbe-asfg-3d4t", true, "Kamal"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}
