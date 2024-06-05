package routes

import (
	"net/http"
)

type Routes struct {
	mux *http.ServeMux
}

func NewRoutes(m *http.ServeMux) *Routes {
	routes := &Routes{
		mux: m,
	}
	routes.registerRoutes()
	return routes
}

func (rs *Routes) registerRoutes() {
	rs.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	rs.mux.HandleFunc("/", HandleIndex)
	rs.mux.HandleFunc("/following", HandleFollowing)
	rs.mux.HandleFunc("/login", HandleLogin)
	rs.mux.HandleFunc("/sample", func(w http.ResponseWriter, r *http.Request) {
		HandleTemplate(w, "body.tmpl", nil)
	})
	//rs.mux.HanldeFunc("/followers", HandleFollowers)
}
