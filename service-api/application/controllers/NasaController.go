package controllers

import (
	"go-young_astrologer/package-core/scalar"
	"go-young_astrologer/service-api/core"
	"go-young_astrologer/service-api/domain/apod"
)

type NasaController struct {
	core *core.Core
}

func NewNasaController(core *core.Core) *NasaController {
	return &NasaController{
		core: core,
	}
}

func (c *NasaController) ApodAction() (interface{}, error) {
	apodService := apod.NewService(c.core.Repositories.Apod)

	return apodService.GetApodList()
}

func (c *NasaController) ApodDateAction(date scalar.Date) (interface{}, error) {
	apodService := apod.NewService(c.core.Repositories.Apod)

	return apodService.GetApodListByDate(date)
}
