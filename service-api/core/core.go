package core

import (
	package_core "go-young_astrologer/package-core"
	"go-young_astrologer/service-api/domain/apod/reposotory"
)

type Core struct {
	*package_core.Core
	Repositories struct {
		Apod reposotory.ApodRepository
	}
}

func NewCore() (*Core, error) {
	baseCore, err := package_core.NewCore()
	if err != nil {
		return nil, err
	}

	return &Core{
		Core: baseCore,
	}, nil
}
