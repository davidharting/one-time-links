package views

import (
	"log"
	"net/http"
)

func showMessageLink(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("showMessage\t")
	props := make(map[string]string)
	link := r.URL.Query().Get("link")
	props["link"] = link
	render(w, "show_message_link", props)

}
