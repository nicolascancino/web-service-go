package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server %v", err)
		panic(err)
	}

}
