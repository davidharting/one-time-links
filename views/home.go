package views

import (
	"fmt"
	"log"
	"net/http"

	"davidharting.com/one-time-links/models"
)

func homeIndex(w http.ResponseWriter, r *http.Request, props map[string]string) {
	log.SetPrefix("view HomeIndex\t")

	err := render(w, "index", props)

	if err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	}
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("view MessageCreate\t")
	log.Println(fmt.Sprintf("Received form submission, message=%v", r.FormValue("message")))

	message := r.FormValue("message")

	props := make(map[string]string)

	link, err := models.EncryptAndSave(message)
	if err != nil {
		log.Printf("Error encrypting and saving message %v\n", err)
		props["alert"] = "Failed to create message"
		homeIndex(w, r, props)
		return
	}

	uri := fmt.Sprintf("%v%v", r.Host, link)
	props["notice"] = fmt.Sprintf("You can share your message with this link %v", uri)
	homeIndex(w, r, props)
}
