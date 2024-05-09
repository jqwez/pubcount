package routes

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandleTemplate(w http.ResponseWriter, templateName string, data interface{}) error {
	fullTemplate := fmt.Sprintf("web/templates/%s", templateName)
	t, err := template.ParseFiles(fullTemplate)
	if err != nil {
		return err
	}
	err = t.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}
