package services

import (
	"context"

	"github.com/Bikes2Road/bikes-compass/internal/core/domain"
	"github.com/Bikes2Road/bikes-compass/internal/core/ports"
	errorBikes "github.com/Bikes2Road/bikes-compass/utils/error"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type placeHolder struct {
	mongoRepository ports.MongoRepository
}

func NewPlaceHolder(mongoRepository ports.MongoRepository) *placeHolder {
	return &placeHolder{
		mongoRepository: mongoRepository,
	}
}

func (s *placeHolder) Execute(ctx context.Context, requestPlaceHolder domain.PlaceHolderRequest) (*domain.PlaceHolderResponseSuccess, *domain.ResponseHttpError) {
	name := requestPlaceHolder.NameByke

	query := bson.M{}

	query["full_name"] = bson.M{"$regex": name, "$options": "i"}

	fields := bson.D{
		{Key: "full_name", Value: 1},
	}

	findOpts := options.Find().SetProjection(fields).SetLimit(5)

	names, err := s.mongoRepository.FindNames(ctx, query, findOpts)

	if err != nil {
		return nil, errorBikes.MapErrorResponse(err.Type, err.Message)
	}

	totalNames := len(names)

	response := &domain.PlaceHolderResponseSuccess{Success: true, Data: names, Total: int64(totalNames)}

	return response, nil
}
