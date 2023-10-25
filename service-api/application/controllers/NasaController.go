package controllers

import (
	package_core "go-young_astrologer/package-core"
	"go-young_astrologer/package-core/scalar"
	"go-young_astrologer/service-api/domain/apod"
	"go-young_astrologer/service-api/infrastructure/db"
)

type NasaController struct {
	core *package_core.Core
}

func NewNasaController(core *package_core.Core) *NasaController {
	return &NasaController{
		core: core,
	}
}

func (c *NasaController) ApodAction() (interface{}, error) {
	apodRepo := db.NewNasaApodRepository(c.core.Engine)
	apodService := apod.NewService(apodRepo)

	return apodService.GetApodList()
}

func (c *NasaController) ApodDateAction(date scalar.Date) (interface{}, error) {
	apodRepo := db.NewNasaApodRepository(c.core.Engine)
	apodService := apod.NewService(apodRepo)

	return apodService.GetApodListByDate(date)
}
