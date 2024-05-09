package main

import (
	"net/http"

	"github.com/jqwez/pubcount/internal"
)

func main() {
	server := internal.NewServer()
	http.ListenAndServe(":3030", server.Mux())
}
