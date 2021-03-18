package main

import (
	"fmt"
	"log"
	"net/http"

	"davidharting.com/one-time-links/views"
)

const PORT = 8080

func main() {
	log.SetPrefix("server: ")
	http.HandleFunc("/", views.Home)

	log.Println(fmt.Sprintf("Server running on port %v 🚀", PORT))
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
