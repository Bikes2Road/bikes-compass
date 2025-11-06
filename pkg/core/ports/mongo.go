package ports

import (
	"context"

	"github.com/Bikes2Road/bikes-compass/pkg/core/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// BikeRepository define la interfaz para el repositorio de bikes
// Esta interfaz permite que los servicios trabajen con el repositorio
// sin depender de la implementación específica de MongoDB
type MongoRepository interface {
	// FindByHash busca una bike por su hash
	FindByHash(ctx context.Context, hash string) (*domain.Bike, error)

	// FindByRef busca una bike por su referencia
	FindByRef(ctx context.Context, ref string) (*domain.Bike, error)

	// FindAll busca todas las bikes que coincidan con el filtro
	FindAll(ctx context.Context, filter bson.M, opts ...options.Lister[options.FindOptions]) ([]*domain.Bike, error)

	// Insert inserta una nueva bike en la colección
	Insert(ctx context.Context, bike *domain.Bike) error

	// InsertMany inserta múltiples bikes en la colección
	InsertMany(ctx context.Context, bikes []*domain.Bike) error

	// UpdateByHash actualiza una bike por su hash
	UpdateByHash(ctx context.Context, hash string, update bson.M) error

	// DeleteByHash elimina una bike por su hash
	DeleteByHash(ctx context.Context, hash string) error

	// Count cuenta el número de bikes que coincidan con el filtro
	Count(ctx context.Context, filter bson.M) (int64, error)

	// SearchByCriteria busca bikes usando criterios de búsqueda flexibles
	SearchByCriteria(ctx context.Context, criteria map[string]interface{}, opts ...options.Lister[options.FindOptions]) ([]*domain.Bike, error)
}

// MongoClient define la interfaz para el cliente de MongoDB
// Esta interfaz permite inyección de dependencias siguiendo arquitectura hexagonal
// y permite que los servicios trabajen con MongoDB sin depender de la implementación específica
type MongoClient interface {
	// Find busca múltiples documentos que coincidan con el filtro
	Find(ctx context.Context, collectionName string, filter bson.M, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error)

	// FindOne busca un solo documento que coincida con el filtro
	FindOne(ctx context.Context, collectionName string, filter bson.M, opts ...options.Lister[options.FindOneOptions]) *mongo.SingleResult

	// InsertOne inserta un documento en la colección
	InsertOne(ctx context.Context, collectionName string, document interface{}) (*mongo.InsertOneResult, error)

	// InsertMany inserta múltiples documentos en la colección
	InsertMany(ctx context.Context, collectionName string, documents []interface{}) (*mongo.InsertManyResult, error)

	// UpdateOne actualiza un documento que coincida con el filtro
	UpdateOne(ctx context.Context, collectionName string, filter bson.M, update bson.M, opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error)

	// UpdateMany actualiza múltiples documentos que coincidan con el filtro
	UpdateMany(ctx context.Context, collectionName string, filter bson.M, update bson.M, opts ...options.Lister[options.UpdateManyOptions]) (*mongo.UpdateResult, error)

	// DeleteOne elimina un documento que coincida con el filtro
	DeleteOne(ctx context.Context, collectionName string, filter bson.M, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error)

	// DeleteMany elimina múltiples documentos que coincidan con el filtro
	DeleteMany(ctx context.Context, collectionName string, filter bson.M, opts ...options.Lister[options.DeleteManyOptions]) (*mongo.DeleteResult, error)

	// CountDocuments cuenta los documentos que coincidan con el filtro
	CountDocuments(ctx context.Context, collectionName string, filter bson.M, opts ...options.Lister[options.CountOptions]) (int64, error)

	// FindOneAndUpdate encuentra y actualiza un documento
	FindOneAndUpdate(ctx context.Context, collectionName string, filter bson.M, update bson.M, opts ...options.Lister[options.FindOneAndUpdateOptions]) *mongo.SingleResult

	// FindOneAndDelete encuentra y elimina un documento
	FindOneAndDelete(ctx context.Context, collectionName string, filter bson.M, opts ...options.Lister[options.FindOneAndDeleteOptions]) *mongo.SingleResult

	// FindOneAndReplace encuentra y reemplaza un documento
	FindOneAndReplace(ctx context.Context, collectionName string, filter bson.M, replacement interface{}, opts ...options.Lister[options.FindOneAndReplaceOptions]) *mongo.SingleResult

	// ReplaceOne reemplaza un documento que coincida con el filtro
	ReplaceOne(ctx context.Context, collectionName string, filter bson.M, replacement interface{}, opts ...options.Lister[options.ReplaceOptions]) (*mongo.UpdateResult, error)

	// GetCollection obtiene una referencia a la colección
	GetCollection(collectionName string) *mongo.Collection

	// Ping verifica la conexión con MongoDB
	Ping(ctx context.Context) error

	// Close cierra la conexión con MongoDB
	Close(ctx context.Context) error
}
