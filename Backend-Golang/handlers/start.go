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
	mux.HandleFunc("/ping", pingPong)
	mux.HandleFunc("/questions", getQuestions)

	fmt.Print("Round table game - Backend Service Started...")

	// starting server
	log.Fatal(http.ListenAndServe("localhost:4040", mux))
}
