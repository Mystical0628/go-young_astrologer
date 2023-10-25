package db

import (
	"go-young_astrologer/package-core/scalar"
	"go-young_astrologer/service-api/domain/apod/entity"
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

func (r *NasaApodRepository) GetApodList() ([]entity.Apod, error) {
	var apodList []entity.Apod

	err := r.engine.Find(&apodList)
	if err != nil {
		return nil, err
	}

	return apodList, nil
}

func (r *NasaApodRepository) GetApodListByDate(date scalar.Date) ([]entity.Apod, error) {
	var apodList []entity.Apod

	err := r.engine.Where("date = ?", date.Format("2006-01-02")).Find(&apodList)
	if err != nil {
		return nil, err
	}

	return apodList, nil
}
