package storage

import (
	"context"
	"io"
	"log/slog"

	"github.com/minio/minio-go/v7"
	"github.com/taupikpirdian/upload-service/domain/repository"
)

type minioStorage struct {
	client *minio.Client
	bucket string
}

func NewMinIOStorage(client *minio.Client, bucket string) repository.StorageRepository {
	return &minioStorage{
		client: client,
		bucket: bucket,
	}
}

func (s *minioStorage) StreamUpload(ctx context.Context, objectName string, reader io.Reader, size int64, contentType string) error {
	opts := minio.PutObjectOptions{
		ContentType: contentType,
		PartSize:    64 * 1024 * 1024,
	}

	info, err := s.client.PutObject(ctx, s.bucket, objectName, reader, size, opts)
	if err != nil {
		return err
	}

	slog.Info("uploaded to minio",
		"object", objectName,
		"size", info.Size,
		"etag", info.ETag,
	)
	return nil
}
