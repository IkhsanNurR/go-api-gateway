package utils

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
)

type FileService struct {
	Client *minio.Client
	Bucket string
}

func NewFileService(client *minio.Client, bucket string) *FileService {
	return &FileService{
		Client: client,
		Bucket: bucket,
	}
}

// Upload a file to the bucket
func (fs *FileService) UploadFile(ctx context.Context, objectName string, file io.Reader, fileSize int64, contentType string) (string, error) {
	opts := minio.PutObjectOptions{
		ContentType: contentType,
	}
	info, err := fs.Client.PutObject(ctx, fs.Bucket, objectName, file, fileSize, opts)
	if err != nil {
		return "", fmt.Errorf("upload failed: %w", err)
	}
	return info.Key, nil
}

// Download a file from the bucket (returns a ReadCloser)
func (fs *FileService) DownloadFile(ctx context.Context, objectName string) (io.ReadCloser, error) {
	object, err := fs.Client.GetObject(ctx, fs.Bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object: %w", err)
	}

	// Optionally do a stat to check existence
	_, err = object.Stat()
	if err != nil {
		return nil, fmt.Errorf("object stat failed: %w", err)
	}

	return object, nil
}

// Delete a file from the bucket
func (fs *FileService) DeleteFile(ctx context.Context, objectName string) error {
	err := fs.Client.RemoveObject(ctx, fs.Bucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}
	return nil
}
