package handlers

import (
	"Backend-Golang/modals"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func pingPong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong")
}

func getQuestions(w http.ResponseWriter, r *http.Request) {
	questions := []modals.Question{
		{"Tallest building in South Asia", "1wfags-qsbe-asfg-3d4t", false, "Raj"},
		{"Tallest building in India", "1wfags-qsbe-asfg-3dst", true, "Kamal"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(questions)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(questions)
	}
}

func getScore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["user_id"])
}
