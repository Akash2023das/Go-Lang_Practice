package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listening on port 9999")
	fmt.Fprintf(w, "Welcome to First Webpage in Go Lang by url %s", r.URL.Path)
	//fmt.Println("Listening on port 9999")

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listening on port 9999")
	fmt.Fprintf(w, "Welcome to about  by url %s", r.URL.Path)
	//fmt.Println("Listening on port 9999")

}
func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/about", aboutHandler)

	//start the server
	http.ListenAndServe(":9999", nil)

}
