package routes

import (
	"net/http"

	"log"
)

type HandleLoginData struct {
	Username   string
	LoginError string
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		HandleGetLogin(w, r)
		return
	}
	if r.Method == "POST" {
		HandlePostLogin(w, r)
	}
}

func HandleGetLogin(w http.ResponseWriter, r *http.Request) {
	data := HandleLoginData{Username: ""}
	loginTemplate(w, data)
}

func HandlePostLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Submission Error", http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	//password := r.FormValue("password")
	data := HandleLoginData{Username: username, LoginError: "Incorrect Details"}

	loginTemplate(w, data)
}

func loginTemplate(w http.ResponseWriter, data HandleLoginData) {
	err := HandleTemplate(w, "login.html", data)
	if err != nil {
		log.Println(err)
	}
}
