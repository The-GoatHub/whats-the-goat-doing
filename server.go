package main

import (
	"context"
	pbGoat "go.thegoathub.dev/whats-the-goat-doing/gen/proto/goat/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
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
	listen, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pbGoat.RegisterGoatServiceServer(server, s)
	reflection.Register(server)

	return server.Serve(listen)
}

func (s *Server) WhatIsDoing(ctx context.Context, in *pbGoat.WhatIsDoingRequest) (*pbGoat.WhatIsDoingResponse, error) {
	action := s.service.GetRandomAction()
	return &pbGoat.WhatIsDoingResponse{
		Action: action,
	}, nil
}
