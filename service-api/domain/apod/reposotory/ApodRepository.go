package reposotory

import (
	"go-young_astrologer/package-core/scalar"
	"go-young_astrologer/service-api/domain/apod/entity"
)

type ApodRepository interface {
	GetApodList() ([]entity.Apod, error)
	GetApodListByDate(date scalar.Date) ([]entity.Apod, error)
}
