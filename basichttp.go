package main

import "fmt"
import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	
	p := r.URL.Path[1:]
	if p == "" {
		fmt.Println("Main Page")
	} else if p == "Chris" {
		fmt.Println("Chris")
	} else {
		fmt.Println("Default")
	}

	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}