package services

import "go-boilerplate/internal/repositories"

type Service struct {
	repository *repositories.Repository
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{
		repository: repository,
	}
}
