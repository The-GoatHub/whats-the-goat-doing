//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"log/slog"
)

var HardcodedSet = wire.NewSet(NewHardcodedRepository, ProvideService)

var InMemorySet = wire.NewSet(GetActionsFromEnv, NewInMemoryRepository, ProvideService)

func ProvideService(repository Repository) (s Service) {
	wire.Build(
		slog.Default,
		NewService,
	)
	return
}

func ProvideHardcodedServer() (s *Server) {
	wire.Build(
		HardcodedSet,
		NewServer,
	)
	return
}

func ProvideInMemoryServer() (s *Server) {
	wire.Build(
		InMemorySet,
		NewServer,
	)
	return
}
