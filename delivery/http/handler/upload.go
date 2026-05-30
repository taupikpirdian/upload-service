package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/taupikpirdian/upload-service/application/usecase"
)

type UploadHandler struct {
	uploadBackup usecase.UploadBackup
}

func NewUploadHandler(uploadBackup usecase.UploadBackup) *UploadHandler {
	return &UploadHandler{
		uploadBackup: uploadBackup,
	}
}

func (h *UploadHandler) Upload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, `{"error":"invalid multipart form"}`, http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, `{"error":"missing form file 'file'"}`, http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := fileHeader.Filename
	if filename == "" {
		filename = fmt.Sprintf("backup-%s.sql.gz", time.Now().UTC().Format("20060102-150405"))
	}

	contentType := fileHeader.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	slog.Info("receiving upload",
		"filename", filename,
		"content_type", contentType,
		"size", fileHeader.Size,
		"remote_addr", r.RemoteAddr,
	)

	resp, err := h.uploadBackup.Execute(r.Context(), filename, file, fileHeader.Size, contentType)
	if err != nil {
		slog.Error("upload failed", "error", err)
		http.Error(w, `{"error":"upload failed"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
