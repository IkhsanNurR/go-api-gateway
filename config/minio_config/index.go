package minioconfig

import (
	"context"
	"log"

	app_config "api-gateway/config/app_config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	Client *minio.Client
	Bucket string
}

func NewMinioClient() *MinioClient {
	endpoint := app_config.EndPoint
	accessKeyID := app_config.AccessKey
	secretAccessKey := app_config.SecretKey
	useSSL := app_config.UseSsl
	bucket := app_config.Bucket

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to MinIO: %v", err)
	}

	// Ensure bucket exists (create if not exist)
	ctx := context.Background()
	exists, errBucketExists := minioClient.BucketExists(ctx, bucket)
	if errBucketExists != nil {
		log.Fatalf("❌ Failed to check bucket: %v", errBucketExists)
	}

	if !exists {
		errCreate := minioClient.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
		if errCreate != nil {
			log.Fatalf("❌ Failed to create bucket: %v", errCreate)
		}
		log.Printf("✅ Bucket %s created.\n", bucket)
	}

	log.Printf("✅ Connected to MinIO at %s (bucket: %s)", endpoint, bucket)

	return &MinioClient{
		Client: minioClient,
		Bucket: bucket,
	}
}
