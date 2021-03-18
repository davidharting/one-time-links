package views

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const templateName string = "index.gohtml"

type templateData struct{}

func Index(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("view /index\t")
	t, err := template.ParseFiles("templates/layout.gohtml", fmt.Sprintf("templates/%v", templateName))
	if err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	} else {
		t.Execute(w, templateData{})
	}

}
