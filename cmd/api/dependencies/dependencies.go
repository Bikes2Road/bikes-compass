package dependencies

import (
	"log"
	"time"

	"github.com/Bikes2Road/bikes-compass/cmd/api/router"
	"github.com/Bikes2Road/bikes-compass/pkg/adapters/cache"
	"github.com/Bikes2Road/bikes-compass/pkg/adapters/http"
	"github.com/Bikes2Road/bikes-compass/pkg/adapters/mongo"
	"github.com/Bikes2Road/bikes-compass/pkg/adapters/r2"
	"github.com/Bikes2Road/bikes-compass/pkg/core"
	"github.com/Bikes2Road/bikes-compass/pkg/core/ports"
	"github.com/Bikes2Road/bikes-compass/utils/env"
	"github.com/gin-gonic/gin"
)

type GetClientMongoFn func(dbName string) (ports.MongoClient, error)
type GetClientR2Fn func() (ports.R2Client, error)
type GetClientCacheFn func(capacity int, ttl time.Duration) ports.CacheClient[string, any]
type NewCacheRepositoryFn func(client ports.CacheClient[string, any]) ports.CacheRepository[string, any]
type NewMongoRepositoryFn func(client ports.MongoClient, collectionName string) ports.MongoRepository
type NewR2RepositoryFn func(client ports.R2Client) ports.R2Repository
type NewApplicationFn func(mongoRepository ports.MongoRepository, r2Repository ports.R2Repository, cacheRepository ports.CacheRepository[string, any]) core.Application
type NewApiHandlerFn func(application core.Application) ports.ApiHandler
type NewRoutesFn func(router *gin.Engine, handlers ports.ApiHandler)

type Wrappers struct {
	newApplication     NewApplicationFn
	getClientMongo     GetClientMongoFn
	getClientR2        GetClientR2Fn
	getClientCache     GetClientCacheFn
	newMongoRepository NewMongoRepositoryFn
	newR2Repository    NewR2RepositoryFn
	newCacheRepository NewCacheRepositoryFn
	newApiHandler      NewApiHandlerFn
	newRoutes          NewRoutesFn
}

func DefaultWrappers() Wrappers {
	return Wrappers{
		newApplication:     core.NewApplication,
		getClientMongo:     mongo.GetClientMongo,
		getClientR2:        r2.GetClientR2,
		getClientCache:     cache.NewCacheClient,
		newMongoRepository: mongo.NewMongoRepository,
		newR2Repository:    r2.NewR2Repository,
		newCacheRepository: cache.NewCacheRepository,
		newApiHandler:      http.NewApiHandler,
		newRoutes:          router.Routers,
	}
}

func InitDependencies(w Wrappers, router *gin.Engine) error {
	log.Println("Iniciando Dependencias...")

	bikesMongoDB := env.GetMongoDBBikes()

	clientMongo, err := w.getClientMongo(bikesMongoDB.DBName)
	if err != nil {
		return err
	}

	mongo.CheckHealth(clientMongo)

	bikesRepository := w.newMongoRepository(clientMongo, bikesMongoDB.Collection)

	clientR2, err := w.getClientR2()

	if err != nil {
		return err
	}

	r2Repository := w.newR2Repository(clientR2)

	cacheClient := w.getClientCache(1000, 90)
	cacheRepository := w.newCacheRepository(cacheClient)

	app := w.newApplication(bikesRepository, r2Repository, cacheRepository)

	handlers := w.newApiHandler(app)

	w.newRoutes(router, handlers)

	return nil
}
