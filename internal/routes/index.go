package routes

import (
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		http.Error(w, "Not Root", http.StatusBadRequest)
		return
	}
	_ = r.Body

	data := GetHandleViewAccountData()
	err := HandleTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
