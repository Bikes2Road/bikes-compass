package mongo

import (
	"context"
	"fmt"

	"github.com/Bikes2Road/bikes-compass/pkg/core/domain"
	"github.com/Bikes2Road/bikes-compass/pkg/core/ports"
	"go.mongodb.org/mongo-driver/v2/bson"
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
func (r *MongoRepository) FindByHash(ctx context.Context, hash string) (*domain.Bike, error) {
	filter := bson.M{"hash_byke": hash}
	result := r.client.FindOne(ctx, r.collectionName, filter)

	if result.Err() != nil {
		return nil, fmt.Errorf("failed to find bike by hash: %w", result.Err())
	}

	var bike domain.Bike
	if err := result.Decode(&bike); err != nil {
		return nil, fmt.Errorf("failed to decode bike: %w", err)
	}

	return &bike, nil
}

// FindByRef busca una bike por su referencia
func (r *MongoRepository) FindByRef(ctx context.Context, ref string) (*domain.Bike, error) {
	filter := bson.M{"ref": ref}
	result := r.client.FindOne(ctx, r.collectionName, filter)

	if result.Err() != nil {
		return nil, fmt.Errorf("failed to find bike by ref: %w", result.Err())
	}

	var bike domain.Bike
	if err := result.Decode(&bike); err != nil {
		return nil, fmt.Errorf("failed to decode bike: %w", err)
	}

	return &bike, nil
}

// FindAll busca todas las bikes que coincidan con el filtro
func (r *MongoRepository) FindAll(ctx context.Context, filter bson.M, opts ...options.Lister[options.FindOptions]) ([]*domain.Bike, error) {
	cursor, err := r.client.Find(ctx, r.collectionName, filter, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to find bikes: %w", err)
	}
	defer cursor.Close(ctx)

	var bikes []*domain.Bike
	if err := cursor.All(ctx, &bikes); err != nil {
		return nil, fmt.Errorf("failed to decode bikes: %w", err)
	}

	return bikes, nil
}

// Insert inserta una nueva bike en la colección
func (r *MongoRepository) Insert(ctx context.Context, bike *domain.Bike) error {
	result, err := r.client.InsertOne(ctx, r.collectionName, bike)
	if err != nil {
		return fmt.Errorf("failed to insert bike: %w", err)
	}

	// Asignar el ID generado al bike
	bike.ID = result.InsertedID
	return nil
}

// InsertMany inserta múltiples bikes en la colección
func (r *MongoRepository) InsertMany(ctx context.Context, bikes []*domain.Bike) error {
	documents := make([]interface{}, len(bikes))
	for i, bike := range bikes {
		documents[i] = bike
	}

	_, err := r.client.InsertMany(ctx, r.collectionName, documents)
	if err != nil {
		return fmt.Errorf("failed to insert bikes: %w", err)
	}

	return nil
}

// UpdateByHash actualiza una bike por su hash
func (r *MongoRepository) UpdateByHash(ctx context.Context, hash string, update bson.M) error {
	filter := bson.M{"hash_byke": hash}

	updateDoc := bson.M{"$set": update}
	result, err := r.client.UpdateOne(ctx, r.collectionName, filter, updateDoc)
	if err != nil {
		return fmt.Errorf("failed to update bike: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("bike with hash %s not found", hash)
	}

	return nil
}

// DeleteByHash elimina una bike por su hash
func (r *MongoRepository) DeleteByHash(ctx context.Context, hash string) error {
	filter := bson.M{"hash_byke": hash}

	result, err := r.client.DeleteOne(ctx, r.collectionName, filter)
	if err != nil {
		return fmt.Errorf("failed to delete bike: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("bike with hash %s not found", hash)
	}

	return nil
}

// Count cuenta el número de bikes que coincidan con el filtro
func (r *MongoRepository) Count(ctx context.Context, filter bson.M) (int64, error) {
	count, err := r.client.CountDocuments(ctx, r.collectionName, filter)
	if err != nil {
		return 0, fmt.Errorf("failed to count bikes: %w", err)
	}

	return count, nil
}

// SearchByCriteria busca bikes usando criterios de búsqueda flexibles
func (r *MongoRepository) SearchByCriteria(ctx context.Context, criteria map[string]interface{}, opts ...options.Lister[options.FindOptions]) ([]*domain.Bike, error) {
	filter := bson.M{}

	// Construir el filtro dinámicamente basado en los criterios
	for key, value := range criteria {
		if value != nil && value != "" {
			filter[key] = value
		}
	}

	return r.FindAll(ctx, filter, opts...)
}
