package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello monish")
}
func sayMonish(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "MONISH PARAMESWARA")
}

// to access this in your computer http://localhost:8060/hello
//       http://localhost:8060/monish
func main() {

	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/monish", sayMonish)

	fmt.Println("The server has started to run.., accces :http://localhost:8060/hello");
	/// always we have to provide the Listen and serve only at the end
	// since once the compiler encounters this line, the control will not go to other line..
	// that is if the handler functions are written after ListenAndServe then that functions will not work..
	http.ListenAndServe(":8060", nil)
}
