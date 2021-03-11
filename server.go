package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Starting server on http://localhost:8080\n")

	//setup static serving
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/r", redirectHandler)

	//startup http server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server! %v", err)
	}

}
