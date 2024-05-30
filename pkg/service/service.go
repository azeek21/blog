package service

import "github.com/azeek21/blog/pkg/repository"

type Service struct {
	Repository *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Repository: repo,
	}
}
