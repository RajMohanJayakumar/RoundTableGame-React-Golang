package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	mux := mux.NewRouter()
	fmt.Print("Round table game - Backend Service Starting...")

	// defining routes
	mux.HandleFunc("/ping", pingPong).Methods(http.MethodGet)
	mux.HandleFunc("/questions", getQuestions).Methods(http.MethodGet)
	mux.HandleFunc("/score/{user_id}", getScore).Methods(http.MethodGet)

	fmt.Print("Round table game - Backend Service Started...")

	// starting server
	log.Fatal(http.ListenAndServe("localhost:4040", mux))
}
