package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	for header, value := range r.Header {
		fmt.Printf("Key: %v \t Value: %v \n", header, value)

	}

	w.Header().Add("TestHeader", "TestValue")
	w.Write([]byte("Hello!"))
}
