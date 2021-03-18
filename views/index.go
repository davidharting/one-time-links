package views

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("view Home\t")
	if r.Method == http.MethodGet {
		log.Println("Handling GET request")
		HomeIndex(w, r)
	} else if r.Method == http.MethodPost {
		log.Println("Handling POST request")
		MessageCreate(w, r)
	} else {
		log.Println(fmt.Sprintf("Unsupported request method %v", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func HomeIndex(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("view HomeIndex\t")
	t, err := template.ParseFiles("templates/layout.gohtml", "templates/index.gohtml")
	if err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	} else {
		t.Execute(w, nil)
	}

}

func MessageCreate(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("view MessageCreate\t")
	log.Println(fmt.Sprintf("Received form submission, message=%v", r.FormValue("message")))
	w.Write([]byte(r.FormValue("message")))
}
