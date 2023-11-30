package main

import (
	"crypto/rand"
	"log/slog"
	"math/big"
)

type Service interface {
	GetRandomAction() string
}

type service struct {
	repository Repository
	logger     *slog.Logger
}

func NewService(repository Repository, logger *slog.Logger) Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (s *service) GetRandomAction() string {
	s.logger.Info("Getting random action")
	actions := s.repository.GetGoatActions()
	randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(actions))))
	action := actions[randomIndex.Int64()]
	s.logger.Info("Got random action", slog.String("action", action))
	return action
}
