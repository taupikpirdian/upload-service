package dto

import "time"

type UploadBackupResponse struct {
	ID         string    `json:"id"`
	Filename   string    `json:"filename"`
	Size       int64     `json:"size"`
	UploadedAt time.Time `json:"uploaded_at"`
}
