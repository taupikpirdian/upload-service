package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/taupikpirdian/upload-service/application/usecase"
	deliveryHttp "github.com/taupikpirdian/upload-service/delivery/http"
	"github.com/taupikpirdian/upload-service/infrastructure/config"
	"github.com/taupikpirdian/upload-service/infrastructure/storage"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	minioClient, err := minio.New(cfg.MinIPEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinIPAccessKey, cfg.MinIPSecretKey, ""),
		Secure: cfg.MinIPUseSSL,
	})
	if err != nil {
		slog.Error("failed to create minio client", "error", err)
		os.Exit(1)
	}

	ensureBucket(minioClient, cfg.MinIPBucket)

	storageRepo := storage.NewMinIOStorage(minioClient, cfg.MinIPBucket)
	uploadBackup := usecase.NewUploadBackup(storageRepo)

	router := deliveryHttp.NewRouter(uploadBackup, cfg.APIKeys)

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  0,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		slog.Info("server starting", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server error", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("server forced to shutdown", "error", err)
	}

	slog.Info("server stopped")
}

func ensureBucket(client *minio.Client, bucket string) {
	ctx := context.Background()
	exists, err := client.BucketExists(ctx, bucket)
	if err != nil {
		slog.Warn("could not check bucket existence", "error", err)
		return
	}
	if !exists {
		if err := client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{}); err != nil {
			slog.Warn("could not create bucket", "bucket", bucket, "error", err)
			return
		}
		slog.Info("created bucket", "bucket", bucket)
	}
}
