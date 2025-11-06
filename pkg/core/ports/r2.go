package ports

import (
	"context"
	"time"
)

type R2Repository interface {
	GetPresignedURL(ctx context.Context, objectKey string, expires time.Duration) (string, error)
	GetBucketName() string
}

type R2Client interface {
	// PresignGetObject genera una URL prefirmada para descargar un objeto del bucket
	PresignGetObject(ctx context.Context, objectKey string, expires time.Duration) (string, error)

	// GetBucketName retorna el nombre del bucket configurado
	GetBucketName() string
}
