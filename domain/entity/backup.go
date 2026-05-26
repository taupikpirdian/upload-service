package entity

import (
	"time"

	"github.com/google/uuid"
)

type Backup struct {
	id          string
	filename    string
	size        int64
	contentType string
	uploadedAt  time.Time
}

func NewBackup(filename, contentType string) *Backup {
	now := time.Now().UTC()
	return &Backup{
		id:          uuid.New().String(),
		filename:    filename,
		contentType: contentType,
		uploadedAt:  now,
	}
}

func ReconstructBackup(id, filename string, size int64, contentType string, uploadedAt time.Time) *Backup {
	return &Backup{
		id:          id,
		filename:    filename,
		size:        size,
		contentType: contentType,
		uploadedAt:  uploadedAt,
	}
}

func (b *Backup) ID() string         { return b.id }
func (b *Backup) Filename() string   { return b.filename }
func (b *Backup) Size() int64        { return b.size }
func (b *Backup) ContentType() string { return b.contentType }
func (b *Backup) UploadedAt() time.Time { return b.uploadedAt }

func (b *Backup) SetSize(size int64) {
	b.size = size
}

func (b *Backup) ObjectPath() string {
	date := b.uploadedAt.Format("2006-01-02")
	return "backups/" + date + "/" + b.id + "-" + b.filename
}
