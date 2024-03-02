package main

import (
	"fmt"
	"net/http"
	"time"

	"rsc.io/quote"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	// http.HandleFunc("/", greet)
	// http.ListenAndServe(":8080", nil)
	fmt.Println(quote.Go())
}