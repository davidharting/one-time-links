package views

import (
	_ "embed"
	"fmt"
	"net/http"
	"text/template"

	"davidharting.com/one-time-links/templates"
)

func render(w http.ResponseWriter, templateName string, props map[string]string) error {
	t, err := template.ParseFS(templates.Files, "layout.gohtml", fmt.Sprintf("%v.gohtml", templateName))
	if err != nil {
		return err
	}
	err = t.Execute(w, props)
	return err
}
