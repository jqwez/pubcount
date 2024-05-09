package routes

import "net/http"

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
	rs.mux.HandleFunc("/", rs.indexRoute)
}

func (rs *Routes) indexRoute(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		http.Error(w, "Not Root", http.StatusBadRequest)
		return
	}
	_ = r.Body
	data := struct{ Title string }{Title: "hi"}
	err := HandleTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
