package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// defining routes
	http.HandleFunc("/ping", pingPong)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:4040", nil))
}

func pingPong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong")
}
