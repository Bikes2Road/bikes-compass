package services

import (
	"context"
	"sync"
	"time"

	"github.com/Bikes2Road/bikes-compass/internal/core/domain"
	"github.com/Bikes2Road/bikes-compass/internal/core/ports"
	errorBikes "github.com/Bikes2Road/bikes-compass/utils/error"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type getAllBikes struct {
	mongoRepository ports.MongoRepository
	r2Repository    ports.R2Repository
	cacheRepository ports.CacheRepository[string, any]
}

func NewGetAllBikes(mongoRepository ports.MongoRepository, r2Repository ports.R2Repository, cacheRepository ports.CacheRepository[string, any]) *getAllBikes {
	return &getAllBikes{
		mongoRepository: mongoRepository,
		r2Repository:    r2Repository,
		cacheRepository: cacheRepository,
	}
}

func (s *getAllBikes) Execute(ctx context.Context, requestByke domain.GetAllBikesRequest, pathRequest string) (*domain.GetAllResponseSuccess, *domain.ResponseHttpError) {

	if cached, ok := s.cacheRepository.GetCached(pathRequest); ok {
		if resp, ok := cached.(*domain.GetAllResponseSuccess); ok {
			return resp, nil
		}
	}

	var query bson.M = bson.M{}
	var fields bson.D = bson.D{}
	var skip int64
	var limit int64
	expireTime := 15 * 60 * time.Second

	query = bson.M{"active": true, "reviewed": true}
	if requestByke.Name != "" {
		query["full_name"] = bson.M{"$regex": requestByke.Name, "$options": "i"}
	}

	if requestByke.Brand != "" {
		query["brand"] = bson.M{"$regex": requestByke.Brand, "$options": "i"}
	}

	skip = (requestByke.Page - 1) * requestByke.Cant
	limit = requestByke.Cant

	// Extract specific fields
	fields = bson.D{
		{Key: "_id", Value: 1},
		{Key: "ref", Value: 1},
		{Key: "hash_byke", Value: 1},
		{Key: "full_name", Value: 1},
		{Key: "year_model", Value: 1},
		{Key: "km", Value: 1},
		{Key: "price", Value: 1},
		{Key: "location", Value: 1},
		{Key: "date_publish", Value: 1},
		{Key: "photos", Value: 1},
	}

	// Para pasar los fields al método FindAll de Mongo, debes crear una opción de proyección y pasarla como opt.
	findOpts := options.Find().SetProjection(fields).SetSkip(skip).SetLimit(limit)

	bikes, err := s.mongoRepository.FindAll(ctx, query, findOpts)

	if err != nil {
		return nil, errorBikes.MapErrorResponse(err.Type, err.Message)
	}

	// Add urls of photos of bikes
	var wg sync.WaitGroup
	for i := range bikes {
		for j := range bikes[i].Photos[0] {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				url, err := s.r2Repository.GetPresignedURL(ctx, bikes[i].Photos[0][j].Key, expireTime)
				if err != nil {
					// If error, set empty string and continue
					bikes[i].Photos[0][j].Url = ""
					return
				}
				bikes[i].Photos[0][j].Url = url
			}(i, j)
		}
	}
	wg.Wait()

	totalBikes := len(bikes)

	response := &domain.GetAllResponseSuccess{Success: true, Data: bikes, Total: int64(totalBikes)}

	s.cacheRepository.SetCached(pathRequest, response)

	return response, nil
}
