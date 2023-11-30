package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	service Service
}

func NewServer(service Service) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/", s.handleGetRandomAction)
	return http.ListenAndServe(":8080", nil)
}

func (s *Server) handleGetRandomAction(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(fmt.Sprintf("The goat is %s\n", s.service.GetRandomAction())))
}
