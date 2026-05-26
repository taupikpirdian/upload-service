package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port           string
	MinIPEndpoint  string
	MinIPAccessKey string
	MinIPSecretKey string
	MinIPBucket    string
	MinIPUseSSL    bool
	APIKeys        []string
}

func Load() *Config {
	return &Config{
		Port:           getEnv("SERVER_PORT", "8080"),
		MinIPEndpoint:  getEnv("MINIO_ENDPOINT", "localhost:9000"),
		MinIPAccessKey: getEnv("MINIO_ACCESS_KEY", "minioadmin"),
		MinIPSecretKey: getEnv("MINIO_SECRET_KEY", "minioadmin"),
		MinIPBucket:    getEnv("MINIO_BUCKET", "backups"),
		MinIPUseSSL:    getEnvBool("MINIO_USE_SSL", false),
		APIKeys:        getEnvSlice("API_KEYS", []string{"test-key"}),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	b, _ := strconv.ParseBool(v)
	return b
}

func getEnvSlice(key string, fallback []string) []string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	parts := strings.Split(v, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			result = append(result, p)
		}
	}
	if len(result) == 0 {
		return fallback
	}
	return result
}
