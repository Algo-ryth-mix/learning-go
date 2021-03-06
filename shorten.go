package main

import (
	"fmt"
	"net/http"
)

// showUrlsHandler [route "/shorten"]
// takes a "long" parameter via POST and generates a new unique short code
// for it that maps to it
// Finally it renders a simple response that show the new mapping in the format
// http://localhost:8080/r?s=key
// x-ref redirectHandler for this route
func shortenHandler(w http.ResponseWriter, r *http.Request) {

	//parsing the POST form
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	//get the unshortened url
	long := r.FormValue("long")

	//get a new code
	counter++
	short := intToLetters(counter)

	//insert the code into the map (note that there is no checking here)
	urlStore[short] = long

	//craft a little bit of a resonse
	fmt.Fprint(w, "<!DOCTYPE html><html lang=\"en\"><body>")
	fmt.Fprintf(w, "New Short Url <a href=http://localhost:8080/r?s=%v>http://localhost:8080/r?s=%v</a>", short, short)
	fmt.Fprint(w, "</body></html>")

}

// redirectHandler [route "/r"]
// takes the parameter "s" via GET that contains a key into the key-value store
// and redirects to the corresponding mapping
// if the key is invalid it renders a simple error message
func redirectHandler(w http.ResponseWriter, r *http.Request) {

	//parse the from (GET this time)
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	//get the key to the map
	key := r.FormValue("s")

	//check if the key exists in the map
	if val, ok := urlStore[key]; ok {

		//if it does redirect the user, using moved permanently
		http.Redirect(w, r, val, http.StatusMovedPermanently)
		return
	}

	//otherwise print that this is an invalid url
	http.Error(w, "Invalid Short URL", http.StatusMisdirectedRequest)
}
