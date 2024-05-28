# Sipinter API

Sipinter API adalah API RESTful untuk mengelola layanan dan otentikasi pengguna. API ini menggunakan MongoDB untuk penyimpanan data dan paket Gorilla Mux untuk routing pada server web Go. API ini menyediakan endpoint untuk operasi CRUD pada layanan serta fungsionalitas registrasi dan login pengguna.

## Powered By

<div style="display: flex; justify-content: left; gap: 20px;">
    <a href="https://go.dev">
        <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="Golang Logo" style="max-width: 100px;">
    </a>
    <a href="https://www.mongodb.com">
        <img src="https://webimages.mongodb.com/_com_assets/cms/kuyjf3vea2hg34taa-horizontal_default_slate_blue.svg" alt="MongoDB Logo" style="max-width: 100px;">
    </a>
</div>

## Fitur

- **Manajemen Layanan**: Membuat, membaca, memperbarui, dan menghapus layanan.
- **Otentikasi Pengguna**: Registrasi dan login pengguna dengan hashing kata sandi yang aman.
- **Middleware**: Middleware untuk otentikasi dan logging untuk mengamankan dan memonitor permintaan API.
- **Integrasi MongoDB**: Menggunakan MongoDB untuk menyimpan dan mengambil data layanan dan pengguna.

## Persiapan

Instruksi berikut akan membantu Anda menyiapkan dan menjalankan proyek ini di mesin lokal Anda.

### Prasyarat

- Go 1.18 atau lebih tinggi
- MongoDB 4.4 atau lebih tinggi

### Instalasi

1. Klon repositori:

   ```bash
   git clone https://github.com/yourusername/sipinter-api.git
   cd sipinter-api
   ```

2. Instal dependensi Go:

   ```bash
   go mod tidy
   ```

3. Jalankan MongoDB di mesin lokal Anda. Pastikan berjalan pada port default `27017`.

4. Jalankan server Go:

   ```bash
   go run main.go
   ```

5. Server akan berjalan di `http://localhost:8080`.

## API Endpoints

### Services

- **Get All Services**

  - **Endpoint**: `GET /api/services`
  - **Description**: Retrieves a list of all services.
  - **Response**: JSON array of service objects.

- **Get Single Service**

  - **Endpoint**: `GET /api/services/{id}`
  - **Description**: Retrieves a single service by its ID.
  - **Response**: JSON object of the service.

- **Create Service**

  - **Endpoint**: `POST /api/services`
  - **Description**: Creates a new service.
  - **Request Body**: JSON object of the service (name, description, location, rating).
  - **Response**: JSON object of the created service.

- **Update Service**

  - **Endpoint**: `PUT /api/services/{id}`
  - **Description**: Updates an existing service by its ID.
  - **Request Body**: JSON object of the updated service (name, description, location, rating).
  - **Response**: JSON object of the updated service.

- **Delete Service**
  - **Endpoint**: `DELETE /api/services/{id}`
  - **Description**: Deletes a service by its ID.
  - **Response**: No content.

### Users

- **Register User**

  - **Endpoint**: `POST /api/users/register`
  - **Description**: Registers a new user.
  - **Request Body**: JSON object with email and password.
  - **Response**: JSON object of the created user.

- **Login User**
  - **Endpoint**: `POST /api/users/login`
  - **Description**: Logs in an existing user.
  - **Request Body**: JSON object with email and password.
  - **Response**: JSON object of the authenticated user (excluding password).

## Middleware

### Middleware Otentikasi

Memastikan bahwa permintaan memiliki token otorisasi yang valid.

### Middleware Logging

Mencatat detail tentang setiap permintaan, termasuk metode, URI, alamat jarak jauh, dan durasi.

## Struktur Proyek

```
sipinter-api/
│
├── handlers/
│   ├── LoginUser.go          # Handler untuk login pengguna
│   ├── RegisterUser.go       # Handler untuk registrasi pengguna
│   └── Service.go            # Handler untuk operasi CRUD layanan
│
├── middlewares/
│   ├── AuthMiddleware.go     # Middleware untuk otentikasi
│   └── LoggingMiddleware.go  # Middleware untuk logging
│
├── models/
│   └── Service.go            # Model untuk layanan dan pengguna
│
├── routes/
│   └── RegisterRoutes.go     # Registrasi rute
│
├── utils/
│   └── ConnectDB.go          # Pengaturan koneksi database
│
├── main.go                   # Entry point dari aplikasi
└── go.mod                    # File modul Go
```

## Pengembang

- **Muhammad Hisyam Kamil** - [hisyam99(https://github.com/hisyam99)
