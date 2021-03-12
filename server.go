package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//TODO find better algorithm to convert int to String sequeneces

func main() {

	//Read backup
	readBackup()

	fmt.Printf("Starting server on http://localhost:8080\n")

	//setup static serving
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/r", redirectHandler)
	http.HandleFunc("/show", showUrlsHandler)

	//startup backup task
	ticker := time.NewTicker(time.Minute * 5)
	quit := make(chan struct{})

	//dispatch goroutine
	go func() {
		for {
			select {

			//check if the event was a ticker event
			case <-ticker.C:
				backupTask()

			//check if the event was a kill event
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	//startup http server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server! %v", err)
	}

	//stop backup task
	close(quit)
}
