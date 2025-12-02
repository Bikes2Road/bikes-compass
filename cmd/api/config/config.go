package config

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
	Server   ServerConfig
	MongoDB  MongoDBConfig
	Cache    CacheConfig
	BucketR2 BucketR2Config
}

type ServerConfig struct {
	Port string
	Host string
	Env  string
}

type MongoDBConfig struct {
	User       string
	Password   string
	AuthSource string
	AppName    string
	Host       string
	Protocol   string
	Uri        string
	Database   string
	Collection string
}

type CacheConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type BucketR2Config struct {
	BucketName      string
	AccountID       string
	TokenValue      string
	AccessKeyID     string
	SecretAccessKey string
}

func Load() (*Config, error) {
	config := &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Host: getEnv("HOST", "0.0.0.0"),
			Env:  getEnv("ENV", "local"),
		},
		MongoDB: MongoDBConfig{
			User:       getEnv("MONGO_USER", ""),
			Password:   getEnv("MONGO_PASSWORD", ""),
			AuthSource: getEnv("MONGO_AUTH_SOURCE", ""),
			AppName:    getEnv("MONGO_APP_NAME", ""),
			Host:       getEnv("MONGO_HOST", ""),
			Protocol:   getEnv("MONGO_PROTOCOL", ""),
			Uri:        getEnv("MONGO_URI", ""),
			Database:   getEnv("MONGO_DATABASE", ""),
			Collection: getEnv("MONGO_COLLECTION", ""),
		},
		Cache: CacheConfig{
			Host:     getEnv("CACHE_HOST", ""),
			Port:     getEnv("CACHE_PORT", ""),
			User:     getEnv("CACHE_USER", ""),
			Password: getEnv("CACHE_PASSWORD", ""),
			Database: getEnv("CACHE_DATABASE", ""),
		},
		BucketR2: BucketR2Config{
			BucketName:      getEnv("BUCKET_NAME", ""),
			AccountID:       getEnv("ACCOUNT_ID", ""),
			TokenValue:      getEnv("TOKEN_VALUE", ""),
			AccessKeyID:     getEnv("ACCESS_KEY_ID", ""),
			SecretAccessKey: getEnv("SECRET_ACCESS_KEY", ""),
		},
	}

	if config.MongoDB.Host == "" || config.MongoDB.Database == "" || config.MongoDB.Collection == "" {
		return nil, errors.New("check env mongo cannot be empty")
	}

	if config.BucketR2.BucketName == "" || config.BucketR2.AccountID == "" || config.BucketR2.TokenValue == "" || config.BucketR2.AccessKeyID == "" || config.BucketR2.SecretAccessKey == "" {
		return nil, errors.New("check env bucket r2 cannot be empty")
	}

	return config, nil

}

// IsDevelopment returns true if running in development mode
func (c *ServerConfig) IsDevelopment() bool {
	return c.Env == "staging" || c.Env == "local"
}

// GetServerAddress returns the server address
func (c *ServerConfig) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
