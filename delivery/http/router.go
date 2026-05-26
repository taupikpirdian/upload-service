package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/taupikpirdian/upload-service/application/usecase"
	"github.com/taupikpirdian/upload-service/delivery/http/handler"
	deliveryMiddleware "github.com/taupikpirdian/upload-service/delivery/http/middleware"
)

func NewRouter(uploadBackup usecase.UploadBackup, apiKeys []string) http.Handler {
	r := chi.NewRouter()

	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	uploadHandler := handler.NewUploadHandler(uploadBackup)

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(deliveryMiddleware.APIKeyAuth(apiKeys))

		r.Post("/backups/upload", uploadHandler.Upload)
	})

	return r
}
