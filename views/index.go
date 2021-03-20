package views

import (
	"fmt"
	"log"
	"net/http"

	"davidharting.com/one-time-links/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("view Home\t")
	if r.Method == http.MethodGet {
		log.Println("Handling GET request")
		homeIndex(w, r, make(map[string]string))
	} else if r.Method == http.MethodPost {
		log.Println("Handling POST request")
		messageCreate(w, r)
	} else {
		log.Println(fmt.Sprintf("Unsupported request method %v", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func homeIndex(w http.ResponseWriter, r *http.Request, props map[string]string) {
	log.SetPrefix("view HomeIndex\t")

	err := render(w, "index", props)

	if err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	}
}

func messageCreate(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("view MessageCreate\t")
	log.Println(fmt.Sprintf("Received form submission, message=%v", r.FormValue("message")))

	message := r.FormValue("message")
	result, err := models.Encrypt(message)

	props := make(
		map[string]string)
	if err != nil {
		props["alert"] = "Failed to create message"
		homeIndex(w, r, props)
		return
	}

	props["notice"] = fmt.Sprintf("Your message has id %v", result.Message.Id)
	homeIndex(w, r, props)
}
