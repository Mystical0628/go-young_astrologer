package apod

import (
	"go-young_astrologer/service-nasa/domain/apod/entity"
	"go-young_astrologer/service-nasa/domain/apod/reposotory"
)

type Service struct {
	repository reposotory.ApodRepository
}

func NewService(repository reposotory.ApodRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) AddApod(entity entity.Apod) error {
	return s.repository.AddApod(entity)
}
