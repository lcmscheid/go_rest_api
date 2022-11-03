package service

import "github.com/lcmscheid/go_rest_api/internal/todo/repository"

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
        repo: r,
    }
}