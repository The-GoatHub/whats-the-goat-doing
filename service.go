package main

import (
	"crypto/rand"
	"math/big"
)

type Service interface {
	GetRandomAction() string
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetRandomAction() string {
	actions := s.repository.GetGoatActions()
	randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(actions))))
	return actions[randomIndex.Int64()]
}
