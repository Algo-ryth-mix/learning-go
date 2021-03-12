package main

import (
	"fmt"
	"net/http"
)

// showUrlsHandler [route "/show"]
// generates a list of all the saved urls and displays them in list
// elements in the html repsonse
func showUrlsHandler(w http.ResponseWriter, _ *http.Request) {

	//start a html list
	fmt.Fprint(w, "<!DOCTYPE html><html lang=\"en\"><body><ul>")

	for key, value := range urlStore {

		//this piece of "code" generates a list entry for each mapping in the form of <short> To <long>
		fmt.Fprintf(w, "<li><a href=\"/r?s=%v\">%v</a> To <a href=%v>%v</a></li>", key, key, value, value)
	}

	//close the html
	fmt.Fprint(w, "</ul></body></html>")
}
