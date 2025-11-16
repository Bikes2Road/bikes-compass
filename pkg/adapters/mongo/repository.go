package mongo

import (
	"context"
	"fmt"

	"github.com/Bikes2Road/bikes-compass/pkg/core/domain"
	"github.com/Bikes2Road/bikes-compass/pkg/core/ports"
	errorBikes "github.com/Bikes2Road/bikes-compass/utils/error"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// BikeRepository implementa el repositorio de bikes usando MongoDB
// Utiliza inyección de dependencias mediante la interfaz MongoClient
type MongoRepository struct {
	client         ports.MongoClient
	collectionName string
}

// NewBikeRepository crea una nueva instancia del repositorio de bikes
// Recibe el cliente MongoDB mediante inyección de dependencias
func NewMongoRepository(client ports.MongoClient, collectionName string) ports.MongoRepository {
	return &MongoRepository{
		client:         client,
		collectionName: collectionName,
	}
}

// FindByHash busca una bike por su hash
func (r *MongoRepository) FindByHash(ctx context.Context, filter bson.M, opts ...options.Lister[options.FindOneOptions]) (*domain.FullBykeResponse, *errorBikes.WrapperError) {
	var byke domain.FullBykeResponse
	err := r.client.FindOne(ctx, r.collectionName, filter, opts...).Decode(&byke)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			newError := fmt.Errorf("failed to find byke by hash: %v", filter["hash_byke"])
			return nil, errorBikes.MapError(errorBikes.ErrorMongoFind, newError)
		} else {
			newError := fmt.Errorf("failed to decode bike: %w", err)
			return nil, errorBikes.MapError(errorBikes.ErrorUnexpected, newError)
		}
	}

	return &byke, nil
}

// FindAll busca todas las bikes que coincidan con el filtro
func (r *MongoRepository) FindAll(ctx context.Context, filter bson.M, opts ...options.Lister[options.FindOptions]) ([]*domain.BykeReponse, *errorBikes.WrapperError) {
	cursor, err := r.client.Find(ctx, r.collectionName, filter, opts...)
	if err != nil {
		//return nil, fmt.Errorf("failed to find bikes: %w", err)
		return nil, errorBikes.MapError(errorBikes.ErrorMongoFindAll, err)
	}
	if cursor == nil {
		return nil, errorBikes.MapError(errorBikes.ErrorUnexpected, nil)
	}
	if cursor.RemainingBatchLength() == 0 {
		return nil, errorBikes.MapError(errorBikes.ErrorBikesNotFound, nil)
	}
	defer cursor.Close(ctx)

	var bikes []*domain.BykeReponse
	if err := cursor.All(ctx, &bikes); err != nil {
		newError := fmt.Errorf("failed to decode bike: %w", err)
		return nil, errorBikes.MapError(errorBikes.ErrorUnexpected, newError)
	}

	for i, byke := range bikes {
		bikes[i].Photos = byke.Photos[:1]
	}

	return bikes, nil
}

// FindAll busca todas las bikes que coincidan con el filtro
func (r *MongoRepository) FindNames(ctx context.Context, filter bson.M, opts ...options.Lister[options.FindOptions]) ([]string, *errorBikes.WrapperError) {
	cursor, err := r.client.Find(ctx, r.collectionName, filter, opts...)
	if err != nil {
		//return nil, fmt.Errorf("failed to find bikes: %w", err)
		return nil, errorBikes.MapError(errorBikes.ErrorMongoFindAll, err)
	}
	if cursor == nil {
		return nil, errorBikes.MapError(errorBikes.ErrorUnexpected, nil)
	}
	if cursor.RemainingBatchLength() == 0 {
		return nil, errorBikes.MapError(errorBikes.ErrorBikesNotFound, nil)
	}
	defer cursor.Close(ctx)

	var bikes []*domain.BykeName
	if err := cursor.All(ctx, &bikes); err != nil {
		newError := fmt.Errorf("failed to decode bike: %w", err)
		return nil, errorBikes.MapError(errorBikes.ErrorUnexpected, newError)
	}

	names := make([]string, len(bikes))
	for i, bike := range bikes {
		if bike != nil {
			names[i] = bike.FullName
		}
	}

	return names, nil
}

// Insert inserta una nueva bike en la colección
func (r *MongoRepository) Insert(ctx context.Context, bike *domain.Bike) *errorBikes.WrapperError {
	result, err := r.client.InsertOne(ctx, r.collectionName, bike)
	if err != nil {
		newError := fmt.Errorf("failed to insert bike: %w", err)
		return errorBikes.MapError(errorBikes.ErrorBadRequest, newError)
	}

	// Asignar el ID generado al bike
	bike.ID = result.InsertedID
	return nil
}

// UpdateByHash actualiza una bike por su hash
func (r *MongoRepository) UpdateByHash(ctx context.Context, hash string, update bson.M) *errorBikes.WrapperError {
	filter := bson.M{"hash_byke": hash}

	updateDoc := bson.M{"$set": update}
	result, err := r.client.UpdateOne(ctx, r.collectionName, filter, updateDoc)
	if err != nil {
		newError := fmt.Errorf("failed to update bike: %w", err)
		return errorBikes.MapError(errorBikes.ErrorUpdateByke, newError)
	}

	if result.MatchedCount == 0 {
		newError := fmt.Errorf("byke with hash %s not found", hash)
		return errorBikes.MapError(errorBikes.ErrorBykeNotFound, newError)
	}

	return nil
}

// DeleteByHash elimina una bike por su hash
func (r *MongoRepository) DeleteByHash(ctx context.Context, hash string) *errorBikes.WrapperError {
	filter := bson.M{"hash_byke": hash}

	result, err := r.client.DeleteOne(ctx, r.collectionName, filter)
	if err != nil {
		newError := fmt.Errorf("failed to delete bike: %w", err)
		return errorBikes.MapError(errorBikes.ErrorDeleteByke, newError)
	}

	if result.DeletedCount == 0 {
		newError := fmt.Errorf("bike with hash %s not found", hash)
		return errorBikes.MapError(errorBikes.ErrorBykeNotFound, newError)
	}

	return nil
}
