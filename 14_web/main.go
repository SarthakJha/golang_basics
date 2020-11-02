package main

import (
	"fmt"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Kasa kai? </h1>")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>this is the / route</h1>")
}

func main() {
	PORT := ":3000"

	// route (/greet), serving function index
	http.HandleFunc("/greet", greet)

	// route(/), serving index
	http.HandleFunc("/", index)

	fmt.Println("listening at port: " + PORT)

	// same as app.listen(PORT,()=> {})
	http.ListenAndServe(PORT, nil)
}
