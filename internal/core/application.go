package core

import (
	"github.com/Bikes2Road/bikes-compass/internal/core/ports"
	"github.com/Bikes2Road/bikes-compass/internal/core/services"
)

type Application struct {
	GetAllBikes ports.GetAllBikes
	GetByke     ports.GetByke
	PlaceHolder ports.PlaceHolder
}

func NewApplication(mongoRepository ports.MongoRepository, r2Repository ports.R2Repository, cacheRepository ports.CacheRepository[string, any]) Application {
	application := Application{
		GetAllBikes: services.NewGetAllBikes(mongoRepository, r2Repository, cacheRepository),
		GetByke:     services.NewGetByke(mongoRepository, r2Repository, cacheRepository),
		PlaceHolder: services.NewPlaceHolder(mongoRepository),
	}

	return application
}
