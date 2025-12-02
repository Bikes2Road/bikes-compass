package r2

import (
	"context"
	"fmt"
	"time"

	configApp "github.com/Bikes2Road/bikes-compass/cmd/api/config"
	"github.com/Bikes2Road/bikes-compass/internal/core/ports"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Client implementa la interfaz R2Client
type NewClientR2 struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	bucketName    string
}

// NewClient crea una nueva instancia del cliente R2
func GetClientR2(r2Credentials configApp.BucketR2Config) (ports.R2Client, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			r2Credentials.AccessKeyID,
			r2Credentials.SecretAccessKey,
			"",
		)),
		config.WithRegion("auto"),
	)
	if err != nil {
		return nil, fmt.Errorf("error loading AWS config: %w", err)
	}

	// Configurar el endpoint de R2
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", r2Credentials.AccountID))
	})

	presignClient := s3.NewPresignClient(client)

	return &NewClientR2{
		client:        client,
		presignClient: presignClient,
		bucketName:    r2Credentials.BucketName,
	}, nil
}

// PresignGetObject genera una URL prefirmada para descargar un objeto
func (c *NewClientR2) PresignGetObject(ctx context.Context, objectKey string, expires time.Duration) (string, error) {
	req, err := c.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = expires
	})
	if err != nil {
		return "", fmt.Errorf("error generating presigned URL: %w", err)
	}
	return req.URL, nil
}

// GetBucketName retorna el nombre del bucket configurado
func (c *NewClientR2) GetBucketName() string {
	return c.bucketName
}
