// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"log/slog"
)

// Injectors from wire.go:

func ProvideService(repository Repository) Service {
	logger := slog.Default()
	mainService := NewService(repository, logger)
	return mainService
}

func ProvideHardcodedServer() *Server {
	repository := NewHardcodedRepository()
	mainService := ProvideService(repository)
	server := NewServer(mainService)
	return server
}

func ProvideInMemoryServer() *Server {
	v := GetActionsFromEnv()
	repository := NewInMemoryRepository(v)
	mainService := ProvideService(repository)
	server := NewServer(mainService)
	return server
}

// wire.go:

var HardcodedSet = wire.NewSet(NewHardcodedRepository, ProvideService)

var InMemorySet = wire.NewSet(GetActionsFromEnv, NewInMemoryRepository, ProvideService)
