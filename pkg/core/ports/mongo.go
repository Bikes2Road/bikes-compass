package ports

import (
	"context"

	"github.com/Bikes2Road/bikes-compass/pkg/core/domain"
	errorBikes "github.com/Bikes2Road/bikes-compass/utils/error"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// BikeRepository define la interfaz para el repositorio de bikes
// Esta interfaz permite que los servicios trabajen con el repositorio
// sin depender de la implementación específica de MongoDB
type MongoRepository interface {
	// FindByHash busca una bike por su hash
	FindByHash(ctx context.Context, filter bson.M, opts ...options.Lister[options.FindOneOptions]) (*domain.FullBykeResponse, *errorBikes.WrapperError)

	// FindAll busca todas las bikes que coincidan con el filtro
	FindAll(ctx context.Context, filter bson.M, opts ...options.Lister[options.FindOptions]) ([]*domain.BykeReponse, *errorBikes.WrapperError)

	// FindNames busca los nombres de las motos que coincidan con los parametros de busqueda
	FindNames(ctx context.Context, filter bson.M, opts ...options.Lister[options.FindOptions]) ([]string, *errorBikes.WrapperError)

	// Insert inserta una nueva bike en la colección
	Insert(ctx context.Context, bike *domain.Bike) *errorBikes.WrapperError

	// UpdateByHash actualiza una bike por su hash
	UpdateByHash(ctx context.Context, hash string, update bson.M) *errorBikes.WrapperError

	// DeleteByHash elimina una bike por su hash
	DeleteByHash(ctx context.Context, hash string) *errorBikes.WrapperError
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

	// UpdateOne actualiza un documento que coincida con el filtro
	UpdateOne(ctx context.Context, collectionName string, filter bson.M, update bson.M, opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error)

	// DeleteOne elimina un documento que coincida con el filtro
	DeleteOne(ctx context.Context, collectionName string, filter bson.M, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error)

	// FindOneAndUpdate encuentra y actualiza un documento
	FindOneAndUpdate(ctx context.Context, collectionName string, filter bson.M, update bson.M, opts ...options.Lister[options.FindOneAndUpdateOptions]) *mongo.SingleResult

	// GetCollection obtiene una referencia a la colección
	GetCollection(collectionName string) *mongo.Collection

	// Ping verifica la conexión con MongoDB
	Ping(ctx context.Context) error

	// Close cierra la conexión con MongoDB
	Close(ctx context.Context) error
}
