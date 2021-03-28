package views

import (
	"log"
	"net/http"

	"github.com/davidharting/one-time-links/models"
)

func showMessage(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("showMessage\t")
	props := make(map[string]string)
	messageId := r.URL.Query().Get("id")
	password := r.URL.Query().Get("password")
	if len(messageId) < 1 {
		props["alert"] = "No message ID in your request"
		render(w, "message_not_found", props)
		return
	}
	if len(password) < 1 {
		props["alert"] = "No password in your request"
		render(w, "message_not_found", props)
		return
	}

	plaintext, err := models.GetMessage(messageId, password)
	if err != nil {
		log.Printf("Error retreiving message %v\n", err)
		render(w, "message_not_found", props)
	} else {
		props["body"] = plaintext
		render(w, "show_message", props)
	}

}
