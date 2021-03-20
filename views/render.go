package views

import (
	"fmt"
	"html/template"
	"net/http"
)

func render(w http.ResponseWriter, templateName string, props map[string]string) error {
	t, err := template.ParseFiles("templates/layout.gohtml", fmt.Sprintf("templates/%v.gohtml", templateName))
	if err != nil {
		return err
	}
	err = t.Execute(w, props)
	return err
}
