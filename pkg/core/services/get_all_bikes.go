package services

import (
	"context"
	"time"

	"github.com/Bikes2Road/bikes-compass/pkg/core/domain"
	"github.com/Bikes2Road/bikes-compass/pkg/core/ports"
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

	//query = bson.M{"sale_status": true}
	query = bson.M{}
	if requestByke.Name != "" {
		query["full_name"] = bson.M{"$regex": requestByke.Name, "$options": "i"}
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

	// Add orls of photos of bikes
	for i, byke := range bikes {
		for j, photo := range byke.Photos[0] {
			bikes[i].Photos[0][j].Url, err = s.r2Repository.GetPresignedURL(ctx, photo.Key, expireTime)
			if err != nil {
				// Remove the bike from the bikes slice if an error occurs generating the photo URL
				bikes = append(bikes[:i], bikes[i+1:]...)
				i--   // decrement i since bikes are now one less and next bike shifts to current index
				break // exit the photo loop for this bike since bike is now removed
			}
		}
	}

	totalBikes := len(bikes)

	response := &domain.GetAllResponseSuccess{Success: true, Data: bikes, Total: int64(totalBikes)}

	s.cacheRepository.SetCached(pathRequest, response)

	return response, nil
}
