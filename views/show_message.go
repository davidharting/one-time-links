package views

import "net/http"

func showMessage(w http.ResponseWriter, r *http.Request) {
	props := make(map[string]string)
	messageId := r.URL.Query().Get("id")
	if len(messageId) < 1 {
		render(w, "message_not_found", props)
		return
	}
	props["body"] = "This would be the secure message"
	render(w, "show_message", props)
}
