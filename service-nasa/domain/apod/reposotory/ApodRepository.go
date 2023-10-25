package reposotory

import "go-young_astrologer/service-nasa/domain/apod/entity"

type ApodRepository interface {
	AddApod(entity entity.Apod) error
}
