package service

import (
	"context"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Ping(ctx context.Context) error {
	return nil
}
