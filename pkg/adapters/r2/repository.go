package r2

import (
	"context"
	"fmt"
	"time"

	"github.com/Bikes2Road/bikes-compass/pkg/core/ports"
	errorBikes "github.com/Bikes2Road/bikes-compass/utils/error"
)

// R2Repository implementa el repositorio para interactuar con objetos en R2
// Utiliza inyección de dependencias mediante la interfaz R2Client
type R2Repository struct {
	client ports.R2Client
}

// NewR2Repository crea una nueva instancia del repositorio R2
// Recibe el cliente R2 mediante inyección de dependencias
func NewR2Repository(client ports.R2Client) ports.R2Repository {
	return &R2Repository{
		client: client,
	}
}

// GetPresignedURL genera una URL prefirmada para descargar un objeto del bucket
func (r *R2Repository) GetPresignedURL(ctx context.Context, objectKey string, expires time.Duration) (string, *errorBikes.WrapperError) {
	if objectKey == "" {
		newError := fmt.Errorf("object key cannot be empty")
		return "", errorBikes.MapError(errorBikes.ErrorR2KeyEmpty, newError)
	}

	key := fmt.Sprintf("n8n_bikes/%s", objectKey)

	url, err := r.client.PresignGetObject(ctx, key, expires)
	if err != nil {
		newError := fmt.Errorf("failed to generate presigned URL: %w", err)
		return "", errorBikes.MapError(errorBikes.ErrorR2Url, newError)
	}

	return url, nil
}

// GetBucketName retorna el nombre del bucket configurado
func (r *R2Repository) GetBucketName() string {
	return r.client.GetBucketName()
}
