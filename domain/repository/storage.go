package repository

import (
	"context"
	"io"
)

type StorageRepository interface {
	StreamUpload(ctx context.Context, objectName string, reader io.Reader, size int64, contentType string) error
}
