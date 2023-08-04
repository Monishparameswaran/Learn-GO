package main

import (
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("static")) // make sure you put all the contents inside the static folder

	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)

	/* how to access

	http://localhost:8080/store.html
	http://localhost:8080/show.html
	*/
}
