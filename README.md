# Project Management API

REST API ini adalah sistem manajemen proyek berbasis Golang menggunakan framework Gin, database PostgreSQL dengan ORM GORM, serta JWT untuk otentikasi.

## Tech Stack
- **Golang** (Gin Framework)
- **PostgreSQL** (Database)
- **GORM** (ORM untuk Golang)
- **JWT** (JSON Web Token untuk otentikasi)

## Instalasi dan Menjalankan Project

### Clone Repository
```bash
git clone https://github.com/fahrillrizal/API-ProjectManagement.git
```

### Konfigurasi Environment
Buat file `.env` berdasarkan template yang ada dan sesuaikan konfigurasi database serta JWT Secret Key:
```env
PGHOST=your_db_host
PGUSER=your_db_user
PGPASSWORD=your_db_password
PGNAME=your_db_name
PGPORT=your_db_port
PORT=your_application_port
JWT_SECRET=your_secret_key
```

### Install Dependencies
```bash
go mod tidy
```

### Menjalankan Aplikasi
```bash
go run main.go
```
Aplikasi akan berjalan di `http://localhost:8080`

## Struktur Folder dan Penjelasan

- **`controllers/`**: Berisi handler untuk menangani HTTP request dan memberikan response.
- **`database/`**: Berisi konfigurasi dan koneksi ke database PostgreSQL.
- **`middleware/`**: Middleware untuk autentikasi
- **`models/`**: Definisi struktur data dan model untuk database menggunakan GORM.
- **`repository/`**: Layer akses database untuk memisahkan logika query dari service.
- **`routes/`**: Menentukan rute dan endpoint API.
- **`services/`**: Berisi logika bisnis aplikasi.
- **`utils/`**: Fungsi utilitas yang jwt(untuk generate dan parse token) dan validation(untuk verifikasi login/register dan hash password).

## Fitur Utama
- Manajemen proyek (CRUD proyek, tugas, dan anggota tim)
- Otentikasi JWT
- Middleware untuk proteksi endpoint
- Dokumentasi API menggunakan Postman

## Dokumentasi API di Postman
Dapat mengakses dokumentasi API melalui Postman dengan mengunjungi link berikut:
[Postman Documentation](https://documenter.getpostman.com/view/22087046/2sAYQiCoX5)