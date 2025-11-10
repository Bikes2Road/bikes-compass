package core

import (
	"github.com/Bikes2Road/bikes-compass/pkg/core/ports"
	"github.com/Bikes2Road/bikes-compass/pkg/core/services"
)

type Application struct {
	GetAllBikes ports.GetAllBikes
}

func NewApplication(mongoRepository ports.MongoRepository, r2Repository ports.R2Repository, cacheRepository ports.CacheRepository[string, any]) Application {
	application := Application{
		GetAllBikes: services.NewGetAllBikes(mongoRepository, r2Repository, cacheRepository),
	}

	return application
}
