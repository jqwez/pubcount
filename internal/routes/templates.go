package routes

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandleTemplate(w http.ResponseWriter, templateName string, data interface{}) error {
	t, err := template.ParseFiles(withLoc("header.tmpl"), withLoc(templateName), withLoc("footer.tmpl"))
	if err != nil {
		http.Error(w, "Templating Error", http.StatusInternalServerError)
		return err
	}
	err = t.ExecuteTemplate(w, templateName, data)
	if err != nil {
		http.Error(w, "Templating Error", http.StatusInternalServerError)
		return err
	}
	return nil
}

func withLoc(templateName string) string {
	return fmt.Sprintf("web/templates/%s", templateName)
}
