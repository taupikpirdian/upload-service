package usecase

import (
	"context"
	"io"

	"github.com/taupikpirdian/upload-service/application/dto"
	"github.com/taupikpirdian/upload-service/domain/entity"
	"github.com/taupikpirdian/upload-service/domain/repository"
)

type UploadBackup interface {
	Execute(ctx context.Context, filename string, reader io.Reader, size int64, contentType string) (*dto.UploadBackupResponse, error)
}

type uploadBackup struct {
	storageRepo repository.StorageRepository
}

func NewUploadBackup(storageRepo repository.StorageRepository) UploadBackup {
	return &uploadBackup{
		storageRepo: storageRepo,
	}
}

func (u *uploadBackup) Execute(ctx context.Context, filename string, reader io.Reader, size int64, contentType string) (*dto.UploadBackupResponse, error) {
	backup := entity.NewBackup(filename, contentType)

	if err := u.storageRepo.StreamUpload(ctx, backup.ObjectPath(), reader, size, backup.ContentType()); err != nil {
		return nil, err
	}

	backup.SetSize(size)

	return &dto.UploadBackupResponse{
		ID:         backup.ID(),
		Filename:   backup.Filename(),
		Size:       backup.Size(),
		UploadedAt: backup.UploadedAt(),
	}, nil
}
