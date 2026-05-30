#!/usr/bin/env bash
# =============================================================================
# Upload Backup API - cURL Examples
# Endpoint: POST /api/v1/backups/upload
# =============================================================================

BASE_URL="http://localhost:8080"
API_KEY="your-api-key"

# -----------------------------------------------------------------------------
# 1. Upload file (filename diambil otomatis dari nama file)
# -----------------------------------------------------------------------------
curl -X POST "${BASE_URL}/api/v1/backups/upload" \
  -H "X-API-Key: ${API_KEY}" \
  --form 'file=@/path/to/backup.sql.gz'

# -----------------------------------------------------------------------------
# 2. Upload file .sql biasa (tanpa kompresi)
# -----------------------------------------------------------------------------
curl -X POST "${BASE_URL}/api/v1/backups/upload" \
  -H "X-API-Key: ${API_KEY}" \
  --form 'file=@/path/to/backup.sql'

# -----------------------------------------------------------------------------
# 3. Upload dengan output verbose (untuk debugging)
# -----------------------------------------------------------------------------
curl -v -X POST "${BASE_URL}/api/v1/backups/upload" \
  -H "X-API-Key: ${API_KEY}" \
  --form 'file=@/path/to/backup.sql.gz'

# -----------------------------------------------------------------------------
# 4. Upload dengan response disimpan ke file
# -----------------------------------------------------------------------------
curl -X POST "${BASE_URL}/api/v1/backups/upload" \
  -H "X-API-Key: ${API_KEY}" \
  --form 'file=@/path/to/backup.sql.gz' \
  -o response.json

# -----------------------------------------------------------------------------
# Contoh Response Sukses (HTTP 201 Created):
# {
#   "id": "550e8400-e29b-41d4-a716-446655440000",
#   "filename": "backup.sql.gz",
#   "size": 204800,
#   "uploaded_at": "2026-05-30T10:15:30Z"
# }
#
# Contoh Response Error - Bad Request (HTTP 400):
# {"error":"invalid multipart form"}
# {"error":"missing form file 'file'"}
#
# Contoh Response Error - Unauthorized (HTTP 401):
# {"error":"missing X-API-Key header"}
# {"error":"invalid api key"}
#
# Contoh Response Error - Upload Gagal (HTTP 500):
# {"error":"upload failed"}
# -----------------------------------------------------------------------------
