package sample

import (
	"context"
	"fmt"
)

type Service interface {
	Hello(ctx context.Context) string
}

type service struct {
	config Config
}

func NewService(config Config) *service {
	return &service{
		config: config,
	}
}

func (s *service) Hello(ctx context.Context) string {
	return fmt.Sprintf("Hello there, %s!", s.config.HelloWho)
}
