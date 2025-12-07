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

type getByke struct {
	mongoRepository ports.MongoRepository
	r2Repository    ports.R2Repository
	cacheRepository ports.CacheRepository[string, any]
}

func NewGetByke(mongoRepository ports.MongoRepository, r2Repository ports.R2Repository, cacheRepository ports.CacheRepository[string, any]) *getByke {
	return &getByke{
		mongoRepository: mongoRepository,
		r2Repository:    r2Repository,
		cacheRepository: cacheRepository,
	}
}

func (s *getByke) Execute(ctx context.Context, requestByke domain.SearchBykeRequest, pathRequest string) (*domain.GetBykeResponseSuccess, *domain.ResponseHttpError) {
	if cached, ok := s.cacheRepository.GetCached(pathRequest); ok {
		if resp, ok := cached.(*domain.GetBykeResponseSuccess); ok {
			return resp, nil
		}
	}

	var query bson.M = bson.M{}
	expireTime := 15 * 60 * time.Second

	//query = bson.M{"sale_status": true}
	query = bson.M{}
	if requestByke.HashByke != "" {
		query["hash_byke"] = requestByke.HashByke
	}

	// Extract specific fields
	// Crear opciones para obtener solo los campos seleccionados usando Projection
	findOpts := options.FindOne().SetProjection(bson.D{
		{Key: "ref", Value: 1},
		{Key: "hash_byke", Value: 1},
		{Key: "full_name", Value: 1},
		{Key: "brand", Value: 1},
		{Key: "model", Value: 1},
		{Key: "cylinder", Value: 1},
		{Key: "engine", Value: 1},
		{Key: "horse_power", Value: 1},
		{Key: "weight", Value: 1},
		{Key: "city_register", Value: 1},
		{Key: "extras", Value: 1},
		{Key: "date_found", Value: 1},
		{Key: "date_soat", Value: 1},
		{Key: "date_tecnico", Value: 1},
		{Key: "page_instagram", Value: 1},
		{Key: "url_post", Value: 1},
		{Key: "year_model", Value: 1},
		{Key: "km", Value: 1},
		{Key: "price", Value: 1},
		{Key: "location", Value: 1},
		{Key: "date_publish", Value: 1},
		{Key: "photos", Value: 1},
		{Key: "torque", Value: 1},
	})

	byke, err := s.mongoRepository.FindByHash(ctx, query, findOpts)
	if err != nil {
		return nil, errorBikes.MapErrorResponse(err.Type, err.Message)
	}

	// Add urls of photos of bike
	var wg sync.WaitGroup
	for i := range byke.Photos {
		for j := range byke.Photos[i] {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				url, err := s.r2Repository.GetPresignedURL(ctx, byke.Photos[i][j].Key, expireTime)
				if err != nil {
					// If error, set empty string and continue
					byke.Photos[i][j].Url = ""
					return
				}
				byke.Photos[i][j].Url = url
			}(i, j)
		}
	}
	wg.Wait()

	response := &domain.GetBykeResponseSuccess{Success: true, Data: byke, Total: 1}

	s.cacheRepository.SetCached(pathRequest, response)

	return response, nil
}
