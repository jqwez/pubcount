package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/jqwez/pubcount/internal"
	"github.com/jqwez/pubcount/internal/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	port := ":3030"
	server := internal.NewServer()
	mainStore := store.NewMongoStore().SetDatabase("pubcount")
	userStore := store.NewUserStore(mainStore)
	user, err := userStore.GetUserByEmail("example@email.com")
	if err != nil {
		log.Println(err)
	}
	log.Println(user)

	log.Println("Running on port", port)
	http.ListenAndServe(port, server.Mux())
}
