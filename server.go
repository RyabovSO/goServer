package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello World")
}

func main() {
	fmt.Println("Listening on port :3000")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)
}