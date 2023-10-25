package commands

import (
	"go-young_astrologer/package-core"
	"go-young_astrologer/service-nasa/domain/apod"
	"go-young_astrologer/service-nasa/infrastructure/db"
	"go-young_astrologer/service-nasa/infrastructure/gateway/nasa"
)

type NasaApodCommand struct {
	core *package_core.Core
}

func NewNasaApodCommand(core *package_core.Core) *NasaApodCommand {
	return &NasaApodCommand{
		core: core,
	}
}

func (c *NasaApodCommand) SaveImageAction() error {
	apodRepo := db.NewNasaApodRepository(c.core.Engine)
	apodGateway := nasa.NewApodGateway(c.core.Config)
	apodService := apod.NewService(apodRepo)

	entity, err := apodGateway.FetchImage(nasa.ApodFetchImageDTO{})
	if err != nil {
		return err
	}

	err = apodService.AddApod(*entity)
	if err != nil {
		return err
	}

	return nil
}
