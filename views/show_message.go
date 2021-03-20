package views

import (
	"log"
	"net/http"

	"davidharting.com/one-time-links/models"
)

func showMessage(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("showMessage\t")
	props := make(map[string]string)
	messageId := r.URL.Query().Get("id")
	if len(messageId) < 1 {
		render(w, "message_not_found", props)
		return
	}

	message, err := models.GetMessage(messageId)
	if err != nil {
		log.Printf("Error retreiving message %v\n", err)
		render(w, "message_not_found", props)
	} else {
		props["id"] = message.Id
		props["body"] = message.Body
		render(w, "show_message", props)
	}

}
