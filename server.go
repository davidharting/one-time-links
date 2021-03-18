package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = 8080

var hitCounter int = 0

func main() {
	log.SetPrefix("server: ")
	http.HandleFunc("/", index)

	log.Println(fmt.Sprintf("Server running on port %v 🚀", PORT))
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	hitCounter++
	fmt.Fprintf(w, "Hello, world %v", hitCounter)
}
