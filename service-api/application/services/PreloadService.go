package services

import (
	"go-young_astrologer/service-api/core"
	"go-young_astrologer/service-api/domain/apod"
	"go-young_astrologer/service-api/infrastructure/db"
	"go-young_astrologer/service-api/infrastructure/memory"
)

type PreloadService struct {
	core *core.Core
}

func NewPreloadService(core *core.Core) *PreloadService {
	return &PreloadService{
		core: core,
	}
}

func (s *PreloadService) Run() error {
	dbApodRepo := db.NewNasaApodRepository(s.core.Engine)
	memoryApodRepo := memory.NewNasaApodRepository()
	apodPreloadService := apod.NewPreloadService(dbApodRepo, memoryApodRepo)
	if err := apodPreloadService.Run(); err != nil {
		return err
	}
	s.core.Repositories.Apod = memoryApodRepo

	return nil
}
