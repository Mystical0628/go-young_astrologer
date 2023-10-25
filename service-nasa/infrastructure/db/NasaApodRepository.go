package db

import (
	"go-young_astrologer/service-nasa/domain/apod/entity"
	"xorm.io/xorm"
)

type NasaApodRepository struct {
	engine *xorm.Engine
}

func NewNasaApodRepository(engine *xorm.Engine) *NasaApodRepository {
	return &NasaApodRepository{
		engine: engine,
	}
}

func (r *NasaApodRepository) AddApod(entity entity.Apod) error {
	_, err := r.engine.InsertOne(entity)
	if err != nil {
		return err
	}

	return nil
}
