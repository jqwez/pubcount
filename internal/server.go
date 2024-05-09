package internal

import (
	"net/http"

	"github.com/jqwez/pubcount/internal/routes"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	mux := http.NewServeMux()
	_ = routes.NewRoutes(mux)
	server := &Server{
		mux: mux,
	}
	return server
}

func (s *Server) Mux() *http.ServeMux {
	return s.mux
}
