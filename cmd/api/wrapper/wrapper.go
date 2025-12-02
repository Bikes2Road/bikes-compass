package wrapper

import (
	"log"
	"time"

	"github.com/Bikes2Road/bikes-compass/cmd/api/config"
	"github.com/Bikes2Road/bikes-compass/internal/adapters/cache"
	"github.com/Bikes2Road/bikes-compass/internal/adapters/http/handlers"
	"github.com/Bikes2Road/bikes-compass/internal/adapters/http/router"
	"github.com/Bikes2Road/bikes-compass/internal/adapters/mongo"
	"github.com/Bikes2Road/bikes-compass/internal/adapters/r2"
	"github.com/Bikes2Road/bikes-compass/internal/core"
	"github.com/Bikes2Road/bikes-compass/internal/core/ports"
)

type GetClientMongoFn func(configMongo config.MongoDBConfig) (ports.MongoClient, error)
type GetClientR2Fn func(r2Credentials config.BucketR2Config) (ports.R2Client, error)
type GetClientCacheFn func(capacity int, ttl time.Duration) ports.CacheClient[string, any]
type NewCacheRepositoryFn func(client ports.CacheClient[string, any]) ports.CacheRepository[string, any]
type NewMongoRepositoryFn func(client ports.MongoClient, collectionName string) ports.MongoRepository
type NewR2RepositoryFn func(client ports.R2Client) ports.R2Repository
type NewApplicationFn func(mongoRepository ports.MongoRepository, r2Repository ports.R2Repository, cacheRepository ports.CacheRepository[string, any]) core.Application
type NewApiHandlerFn func(application core.Application) ports.ApiHandler
type NewRoutesFn func(handlers ports.ApiHandler) ports.Router

type Wrapper struct {
	Config             *config.Config
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

func DefaultWrapper() *Wrapper {
	return &Wrapper{
		newApplication:     core.NewApplication,
		getClientMongo:     mongo.GetClientMongo,
		getClientR2:        r2.GetClientR2,
		getClientCache:     cache.NewCacheClient,
		newMongoRepository: mongo.NewMongoRepository,
		newR2Repository:    r2.NewR2Repository,
		newCacheRepository: cache.NewCacheRepository,
		newApiHandler:      handlers.NewApiHandler,
		newRoutes:          router.NewRouter,
	}
}

type App struct {
	Config          *config.Config
	MongoRepository ports.MongoRepository
	R2Repository    ports.R2Repository
	CacheRepository ports.CacheRepository[string, any]
	Application     core.Application
	ApiHandler      ports.ApiHandler
	Router          ports.Router
}

func NewApp(w *Wrapper, cfg *config.Config) (*App, error) {
	app := &App{
		Config: cfg,
	}

	log.Println("Iniciando Dependencias...")

	clientMongo, err := w.getClientMongo(cfg.MongoDB)
	if err != nil {
		return nil, err
	}

	mongo.CheckHealth(clientMongo)

	app.MongoRepository = w.newMongoRepository(clientMongo, cfg.MongoDB.Collection)

	clientR2, err := w.getClientR2(cfg.BucketR2)

	if err != nil {
		log.Panicf("error loading AWS config: %v", err)
	}

	app.R2Repository = w.newR2Repository(clientR2)

	cacheClient := w.getClientCache(1000, 90)
	app.CacheRepository = w.newCacheRepository(cacheClient)

	app.Application = w.newApplication(app.MongoRepository, app.R2Repository, app.CacheRepository)

	app.ApiHandler = w.newApiHandler(app.Application)

	app.Router = w.newRoutes(app.ApiHandler)

	return app, nil
}
