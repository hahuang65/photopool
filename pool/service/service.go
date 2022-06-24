package service

import (
	"git.sr.ht/~hwrd/photopool/pool"
	"git.sr.ht/~hwrd/photopool/pool/repository"
	"github.com/google/uuid"
)

type Service struct {
	repository *repository.Repository
}

func New(r repository.Repository) *Service {
	return &Service{repository: &r}
}

func (s Service) CreateWithName(name string) (pool.Pool, error) {
	p := pool.New(name)
	return (*s.repository).Create(&p)
}

func (s Service) GetByID(id uuid.UUID) (pool.Pool, error) {
	return (*s.repository).Get(id)
}
