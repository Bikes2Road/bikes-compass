package env

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type R2Credentials struct {
	AccountId       string
	TokenValue      string
	AccessKeyId     string
	SecretAccessKey string
	BucketName      string
}

type MongoDBFields struct {
	User       string
	Password   string
	AuthSource string
	Protocol   string
	Host       string
	AppName    string
	Uri        string
}

type RedisFields struct {
	Endpoint string
	Password string
	DataBase int
}

type DBBikesMongo struct {
	DBName     string
	Collection string
}

// Check if env file is enabled
func LoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

// Extract type of DB to Connect ATLAS or LOCAL
func GetEnvironment() string {
	LoadEnvFile()
	return os.Getenv("ENVIRONMENT")
}

// Extract Credentials to Connect with MongoDB
func GetMongoDBCredentials(endpoint string) MongoDBFields {
	var mongoFields MongoDBFields
	endpoint = strings.ToUpper(endpoint)

	LoadEnvFile()

	mongoFields.User = os.Getenv(fmt.Sprintf("MONGO_%s_USER", endpoint))
	mongoFields.Password = os.Getenv(fmt.Sprintf("MONGO_%s_PW", endpoint))
	mongoFields.AuthSource = os.Getenv(fmt.Sprintf("MONGO_%s_AUTHSOURCE", endpoint))
	mongoFields.Host = os.Getenv(fmt.Sprintf("MONGO_%s_HOST", endpoint))
	mongoFields.Protocol = os.Getenv(fmt.Sprintf("MONGO_%s_PROTOCOL", endpoint))
	mongoFields.AppName = os.Getenv(fmt.Sprintf("MONGO_%s_APPNAME", endpoint))
	mongoFields.Uri = fmt.Sprintf("%s://%s", mongoFields.Protocol, mongoFields.Host)

	return mongoFields
}

// Extract credentials of Redis
func GetRedisCredentials() RedisFields {
	var redisFields RedisFields
	var err error

	LoadEnvFile()

	redisFields.Endpoint = os.Getenv("REDIS_ENDPOINT")
	redisFields.Password = os.Getenv("REDIS_PASSWORD")
	redisFields.DataBase, err = strconv.Atoi(os.Getenv("REDIS_DATABASE"))

	if err != nil {
		log.Fatal(err)
	}

	return redisFields
}

// Extract Credentials to Create client of R2
func GetR2Credentials() R2Credentials {

	var r2Credentials R2Credentials

	LoadEnvFile()

	r2Credentials.AccountId = os.Getenv("ACCOUNT_ID")
	r2Credentials.TokenValue = os.Getenv("TOKEN_VALUE")
	r2Credentials.AccessKeyId = os.Getenv("ACCESS_KEY_ID")
	r2Credentials.SecretAccessKey = os.Getenv("SECRET_ACCESS_KEY")
	r2Credentials.BucketName = os.Getenv("BUCKET_NAME")

	return r2Credentials
}

func GetMongoDBBikes() DBBikesMongo {
	var dbBikesMongo DBBikesMongo

	LoadEnvFile()

	dbBikesMongo.DBName = os.Getenv("BIKES_MONGODB_NAME")
	dbBikesMongo.Collection = os.Getenv("BIKES_MONGODB_COLLECTION")

	return dbBikesMongo
}
