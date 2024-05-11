package main

import (
	"log"
	"net/http"

	"github.com/jqwez/pubcount/internal"
)

func main() {
	port := ":3030"
	server := internal.NewServer()
	log.Println("Running on port", port)
	http.ListenAndServe(port, server.Mux())
}
