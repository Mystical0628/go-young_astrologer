package reposotory

import (
	"go-young_astrologer/service-api/domain/apod/entity"
)

type TargetApodRepository interface {
	SetApodList(entity []entity.Apod) error
}
