#!/usr/bin/env bash
# =============================================================================
# Upload Backup API - cURL Examples
# Endpoint: POST /api/v1/backups/upload
# =============================================================================

BASE_URL="http://localhost:8080"
API_KEY="your-api-key"

# -----------------------------------------------------------------------------
# 1. Upload file dengan filename otomatis (format: backup-YYYYMMDD-HHMMSS.sql.gz)
# -----------------------------------------------------------------------------
curl -X POST "${BASE_URL}/api/v1/backups/upload" \
  -H "X-API-Key: ${API_KEY}" \
  -H "Content-Type: application/octet-stream" \
  --data-binary @/path/to/backup.sql.gz

# -----------------------------------------------------------------------------
# 2. Upload file dengan custom filename via query parameter
# -----------------------------------------------------------------------------
curl -X POST "${BASE_URL}/api/v1/backups/upload?filename=mydb-backup-2026-05-30.sql.gz" \
  -H "X-API-Key: ${API_KEY}" \
  -H "Content-Type: application/octet-stream" \
  --data-binary @/path/to/backup.sql.gz

# -----------------------------------------------------------------------------
# 3. Upload file .sql biasa (tanpa kompresi)
# -----------------------------------------------------------------------------
curl -X POST "${BASE_URL}/api/v1/backups/upload?filename=mydb-backup.sql" \
  -H "X-API-Key: ${API_KEY}" \
  -H "Content-Type: text/plain" \
  --data-binary @/path/to/backup.sql

# -----------------------------------------------------------------------------
# 4. Upload dengan Content-Length eksplisit (untuk file besar / streaming)
# -----------------------------------------------------------------------------
FILE="/path/to/backup.sql.gz"
FILE_SIZE=$(stat -f%z "${FILE}" 2>/dev/null || stat -c%s "${FILE}")

curl -X POST "${BASE_URL}/api/v1/backups/upload?filename=large-backup.sql.gz" \
  -H "X-API-Key: ${API_KEY}" \
  -H "Content-Type: application/octet-stream" \
  -H "Content-Length: ${FILE_SIZE}" \
  --data-binary @"${FILE}"

# -----------------------------------------------------------------------------
# 5. Upload dengan output verbose (untuk debugging)
# -----------------------------------------------------------------------------
curl -v -X POST "${BASE_URL}/api/v1/backups/upload" \
  -H "X-API-Key: ${API_KEY}" \
  -H "Content-Type: application/octet-stream" \
  --data-binary @/path/to/backup.sql.gz

# -----------------------------------------------------------------------------
# 6. Upload dengan response disimpan ke file
# -----------------------------------------------------------------------------
curl -X POST "${BASE_URL}/api/v1/backups/upload" \
  -H "X-API-Key: ${API_KEY}" \
  -H "Content-Type: application/octet-stream" \
  --data-binary @/path/to/backup.sql.gz \
  -o response.json

# -----------------------------------------------------------------------------
# Contoh Response Sukses (HTTP 201 Created):
# {
#   "id": "550e8400-e29b-41d4-a716-446655440000",
#   "filename": "backup-20260530-101530.sql.gz",
#   "size": 204800,
#   "uploaded_at": "2026-05-30T10:15:30Z"
# }
#
# Contoh Response Error - Unauthorized (HTTP 401):
# {"error":"missing X-API-Key header"}
# {"error":"invalid api key"}
#
# Contoh Response Error - Upload Gagal (HTTP 500):
# {"error":"upload failed"}
# -----------------------------------------------------------------------------
