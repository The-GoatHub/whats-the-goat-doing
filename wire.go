//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func ProvideRepository() (r Repository) {
	wire.Build(
		NewRepository,
	)
	return
}

func ProvideService() (s Service) {
	wire.Build(
		ProvideRepository,
		NewService,
	)
	return
}

func ProvideServer() (s *Server) {
	wire.Build(
		ProvideService,
		NewServer,
	)
	return
}
