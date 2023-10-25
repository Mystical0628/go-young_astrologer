package memory

import (
	"go-young_astrologer/package-core/scalar"
	"go-young_astrologer/service-api/domain/apod/entity"
)

type NasaApodRepository struct {
	entities []entity.Apod
}

func NewNasaApodRepository() *NasaApodRepository {
	return &NasaApodRepository{}
}

func (r *NasaApodRepository) SetApodList(entities []entity.Apod) error {
	r.entities = entities
	return nil
}

func (r *NasaApodRepository) GetApodList() ([]entity.Apod, error) {
	return r.entities, nil
}

func (r *NasaApodRepository) GetApodListByDate(date scalar.Date) ([]entity.Apod, error) {
	filteredEntities := make([]entity.Apod, 0, len(r.entities))
	for _, apod := range r.entities {
		if apod.Date.Equal(date) {
			filteredEntities = append(filteredEntities, apod)
		}
	}

	return filteredEntities, nil
}
