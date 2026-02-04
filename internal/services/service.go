package services

import "triple-s/internal/repository"

type Service struct {
	store *repository.Store
}

func NewService(store *repository.Store) *Service {
	return &Service{store: store}
}
