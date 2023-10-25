package apod

import (
	"go-young_astrologer/package-core/scalar"
	"go-young_astrologer/service-api/domain/apod/entity"
	"go-young_astrologer/service-api/domain/apod/reposotory"
)

type Service struct {
	repository reposotory.ApodRepository
}

func NewService(repository reposotory.ApodRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetApodList() ([]entity.Apod, error) {
	return s.repository.GetApodList()
}

func (s *Service) GetApodListByDate(date scalar.Date) ([]entity.Apod, error) {
	return s.repository.GetApodListByDate(date)
}
