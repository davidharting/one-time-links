package views

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

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

	relative_link, err := models.EncryptAndSave(message)
	if err != nil {
		log.Printf("Error encrypting and saving message %v\n", err)
		props["alert"] = "Failed to create message"
		homeIndex(w, r, props)
		return
	}

	link := fmt.Sprintf("%v%v", r.Host, relative_link)
	params := url.Values{"link": {link}}
	redirect_to := fmt.Sprintf("/link/?%v", params.Encode())
	http.Redirect(w, r, redirect_to, http.StatusSeeOther)
}
