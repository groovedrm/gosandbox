package main

import "fmt"
import "net/http"


func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Main Page")
}

func chris_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chris's Page")
}


func main() {
	
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/Chris", chris_handler)
	http.ListenAndServe(":8080", nil)
}