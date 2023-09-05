package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Pong")
	})

	http.ListenAndServe("localhost:4040", nil)
}
