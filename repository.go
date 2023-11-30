package main

import (
	"os"
	"strings"
)

type Repository interface {
	GetGoatActions() []string
}

type hardcodedRepository struct{}

func NewHardcodedRepository() Repository {
	return &hardcodedRepository{}
}

func (r *hardcodedRepository) GetGoatActions() []string {
	return []string{
		"eating grass",
		"riding a scooter",
		"drinking a coffee",
	}
}

type inMemoryRepository struct{}

func GetActionsFromEnv() []string {
	return strings.Split(os.Getenv("GOAT_ACTIONS"), ",")
}

func NewInMemoryRepository(actions []string) Repository {
	return &inMemoryRepository{}
}

func (r *inMemoryRepository) GetGoatActions() []string {
	return []string{
		"eating a pizza",
		"using a computer",
		"drinking water",
	}
}
