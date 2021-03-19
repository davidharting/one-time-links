package views

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type LayoutProps struct {
	alert  string
	notice string
}

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

	err := render(w, "index", LayoutProps{})

	if err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	}

}

func MessageCreate(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("view MessageCreate\t")
	log.Println(fmt.Sprintf("Received form submission, message=%v", r.FormValue("message")))
	w.Write([]byte(r.FormValue("message")))
}

type TemplateData struct {
	Alert     string
	HasNotice bool
	HasAlert  bool
	Notice    string
}

func render(w http.ResponseWriter, templateName string, props LayoutProps) error {
	data := getTemplateData(props)
	t, err := template.ParseFiles("templates/layout.gohtml", fmt.Sprintf("templates/%v.gohtml", templateName))
	if err != nil {
		return err
	}
	err = t.Execute(w, data)
	return err
}

func getTemplateData(props LayoutProps) TemplateData {
	hasNotice := len(props.notice) > 0
	hasAlert := len(props.alert) > 0
	return TemplateData{
		Alert:     fmt.Sprintf(props.alert),
		HasAlert:  hasAlert,
		HasNotice: hasNotice,
		Notice:    props.notice,
	}
}
