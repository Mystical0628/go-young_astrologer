package apod

import (
	"go-young_astrologer/service-api/domain/apod/reposotory"
)

type PreloadService struct {
	sourceRepository reposotory.ApodRepository
	targetRepository reposotory.TargetApodRepository
}

func NewPreloadService(sourceRepository reposotory.ApodRepository, targetRepository reposotory.TargetApodRepository) *PreloadService {
	return &PreloadService{
		sourceRepository: sourceRepository,
		targetRepository: targetRepository,
	}
}

func (s *PreloadService) Run() error {
	list, err := s.sourceRepository.GetApodList()
	if err != nil {
		return err
	}

	err = s.targetRepository.SetApodList(list)
	if err != nil {
		return err
	}

	return nil
}
