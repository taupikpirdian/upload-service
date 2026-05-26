# Project: Backup Upload Service

## Objective

Membuat service backend menggunakan Golang untuk menerima upload file backup database (misalnya `.sql`, `.sql.gz`, `.zip`) dari server lain melalui HTTP API menggunakan `curl`.

File backup akan disimpan ke MinIO.

Tambahkan API Key untuk menggunakan API ini.

Service harus aman, ringan, dan support file besar (hingga beberapa GB).

Struktur folder bisa cek di skills

Contoh use case:

Server database melakukan:

```bash
mysqldump mydb | gzip | curl \
-X POST https://backup-service.domain.com/api/v1/backups/upload \
-H "Authorization: Bearer YOUR_SECRET_TOKEN" \
--data-binary @-