package server

import (
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	log.Println("shortlink start on port: " + port)
	return s.httpServer.ListenAndServe()
}
