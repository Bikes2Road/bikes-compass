package mongo

import (
	"context"
	"fmt"

	"github.com/Bikes2Road/bikes-compass/pkg/core/ports"
	"github.com/Bikes2Road/bikes-compass/utils/env"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Client implementa la interfaz MongoClient
type NewClientMongo struct {
	client   *mongo.Client
	database *mongo.Database
	dbName   string
}

// NewClient crea una nueva instancia del cliente MongoDB
func GetClientMongo(dbName string) (ports.MongoClient, error) {
	client, err := Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	database := client.Database(dbName)

	return &NewClientMongo{
		client:   client,
		database: database,
		dbName:   dbName,
	}, nil
}

// Connect establece la conexión con MongoDB
func Connect() (*mongo.Client, error) {
	var clientOpts *options.ClientOptions

	// Extract Values from env file
	environment := env.GetEnvironment()
	credentials := env.GetMongoDBCredentials(environment)

	// Make credentials to connect DB
	credential := options.Credential{
		AuthSource: credentials.AuthSource,
		Username:   credentials.User,
		Password:   credentials.Password,
	}

	// Make Connection to DB on ATLAS or LOCAL
	if environment == "ATLAS" {
		clientOpts = options.Client().ApplyURI(credentials.Uri).SetAuth(credential).SetRetryWrites(true).SetAppName(credentials.AppName)
	} else if environment == "LOCAL" {
		clientOpts = options.Client().ApplyURI(credentials.Uri).SetAuth(credential)
	}

	client, err := mongo.Connect(clientOpts)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	return client, nil
}

// CheckHealth verifica el estado de la conexión con MongoDB
func CheckHealth(client ports.MongoClient) error {
	err := client.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("error trying to ping MongoDB: %w", err)
	}
	fmt.Println("[MongoDB] Connection to MongoDB Success")
	return nil
}

// GetCollection obtiene una referencia a la colección
func (c *NewClientMongo) GetCollection(collectionName string) *mongo.Collection {
	return c.database.Collection(collectionName)
}

// Find busca múltiples documentos
func (c *NewClientMongo) Find(ctx context.Context, collectionName string, filter bson.M, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error) {
	collection := c.GetCollection(collectionName)
	return collection.Find(ctx, filter, opts...)
}

// FindOne busca un solo documento
func (c *NewClientMongo) FindOne(ctx context.Context, collectionName string, filter bson.M, opts ...options.Lister[options.FindOneOptions]) *mongo.SingleResult {
	collection := c.GetCollection(collectionName)
	return collection.FindOne(ctx, filter, opts...)
}

// InsertOne inserta un documento
func (c *NewClientMongo) InsertOne(ctx context.Context, collectionName string, document interface{}) (*mongo.InsertOneResult, error) {
	collection := c.GetCollection(collectionName)
	return collection.InsertOne(ctx, document)
}

// InsertMany inserta múltiples documentos
func (c *NewClientMongo) InsertMany(ctx context.Context, collectionName string, documents []interface{}) (*mongo.InsertManyResult, error) {
	collection := c.GetCollection(collectionName)
	return collection.InsertMany(ctx, documents)
}

// UpdateOne actualiza un documento
func (c *NewClientMongo) UpdateOne(ctx context.Context, collectionName string, filter bson.M, update bson.M, opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error) {
	collection := c.GetCollection(collectionName)
	return collection.UpdateOne(ctx, filter, update, opts...)
}

// UpdateMany actualiza múltiples documentos
func (c *NewClientMongo) UpdateMany(ctx context.Context, collectionName string, filter bson.M, update bson.M, opts ...options.Lister[options.UpdateManyOptions]) (*mongo.UpdateResult, error) {
	collection := c.GetCollection(collectionName)
	return collection.UpdateMany(ctx, filter, update, opts...)
}

// DeleteOne elimina un documento
func (c *NewClientMongo) DeleteOne(ctx context.Context, collectionName string, filter bson.M, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error) {
	collection := c.GetCollection(collectionName)
	return collection.DeleteOne(ctx, filter, opts...)
}

// DeleteMany elimina múltiples documentos
func (c *NewClientMongo) DeleteMany(ctx context.Context, collectionName string, filter bson.M, opts ...options.Lister[options.DeleteManyOptions]) (*mongo.DeleteResult, error) {
	collection := c.GetCollection(collectionName)
	return collection.DeleteMany(ctx, filter, opts...)
}

// CountDocuments cuenta los documentos
func (c *NewClientMongo) CountDocuments(ctx context.Context, collectionName string, filter bson.M, opts ...options.Lister[options.CountOptions]) (int64, error) {
	collection := c.GetCollection(collectionName)
	return collection.CountDocuments(ctx, filter, opts...)
}

// FindOneAndUpdate encuentra y actualiza un documento
func (c *NewClientMongo) FindOneAndUpdate(ctx context.Context, collectionName string, filter bson.M, update bson.M, opts ...options.Lister[options.FindOneAndUpdateOptions]) *mongo.SingleResult {
	collection := c.GetCollection(collectionName)
	return collection.FindOneAndUpdate(ctx, filter, update, opts...)
}

// FindOneAndDelete encuentra y elimina un documento
func (c *NewClientMongo) FindOneAndDelete(ctx context.Context, collectionName string, filter bson.M, opts ...options.Lister[options.FindOneAndDeleteOptions]) *mongo.SingleResult {
	collection := c.GetCollection(collectionName)
	return collection.FindOneAndDelete(ctx, filter, opts...)
}

// FindOneAndReplace encuentra y reemplaza un documento
func (c *NewClientMongo) FindOneAndReplace(ctx context.Context, collectionName string, filter bson.M, replacement interface{}, opts ...options.Lister[options.FindOneAndReplaceOptions]) *mongo.SingleResult {
	collection := c.GetCollection(collectionName)
	return collection.FindOneAndReplace(ctx, filter, replacement, opts...)
}

// ReplaceOne reemplaza un documento
func (c *NewClientMongo) ReplaceOne(ctx context.Context, collectionName string, filter bson.M, replacement interface{}, opts ...options.Lister[options.ReplaceOptions]) (*mongo.UpdateResult, error) {
	collection := c.GetCollection(collectionName)
	return collection.ReplaceOne(ctx, filter, replacement, opts...)
}

// Ping verifica la conexión con MongoDB
func (c *NewClientMongo) Ping(ctx context.Context) error {
	return c.client.Ping(ctx, nil)
}

// Close cierra la conexión con MongoDB
func (c *NewClientMongo) Close(ctx context.Context) error {
	return c.client.Disconnect(ctx)
}
