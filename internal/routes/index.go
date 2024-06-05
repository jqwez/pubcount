package routes

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/jqwez/pubcount/web/pages"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		http.Error(w, "Not Root", http.StatusBadRequest)
		http.Redirect(w, r, "/404", http.StatusNotFound)
		return
	}
	_ = r.Body

	Render(pages.IndexPage()).ServeHTTP(w, r)
}

func Render(c templ.Component) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Render(r.Context(), w)
	})
}
