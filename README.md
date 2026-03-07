# 📄 CV Builder API

RESTful API backend untuk aplikasi **CV Builder** — membantu user membuat dan mengelola data CV/Resume secara digital. Dibangun dengan **Go (Golang)** menggunakan arsitektur **Clean Architecture**.

---

## ⚡ Tech Stack

| Teknologi | Deskripsi |
|---|---|
| **Go 1.25** | Bahasa pemrograman utama |
| **Chi v5** | HTTP router yang ringan & kompatibel `net/http` |
| **GORM** | ORM untuk interaksi database |
| **PostgreSQL** | Database relasional |
| **JWT (HS256)** | Autentikasi berbasis token (berlaku 24 jam) |
| **bcrypt** | Hashing password |
| **godotenv** | Manajemen environment variable |

---

## 📁 Struktur Folder

```
cv_builder_be/
├── main.go                  # Entry point & konfigurasi routing
├── .env                     # Environment variables (tidak di-commit)
├── go.mod / go.sum          # Go module dependencies
│
├── config/
│   ├── config.go            # Struct & loader konfigurasi dari .env
│   └── database.go          # Koneksi PostgreSQL & auto-migrate
│
├── internal/
│   ├── models/              # Definisi struct / skema database
│   │   ├── user.go          # User (email, password, relasi ke semua data CV)
│   │   ├── profiles.go      # Profile (nama, telepon, alamat, foto, summary)
│   │   ├── work_experiences.go  # Pengalaman kerja
│   │   ├── education.go     # Riwayat pendidikan
│   │   ├── skills.go        # Keahlian (nama & level)
│   │   └── projects.go      # Proyek portfolio
│   │
│   ├── repository/          # Layer akses database (GORM queries)
│   │   ├── user_repository.go
│   │   ├── profile_repository.go
│   │   └── cv_repository.go     # (placeholder untuk fitur mendatang)
│   │
│   ├── usecase/             # Business logic
│   │   ├── user_usecase.go      # Register & Login logic
│   │   └── profile_usecase.go   # Get & Upsert profile logic
│   │
│   └── delivery/            # HTTP handlers (controller)
│       ├── user_handler.go      # Handler register & login
│       └── profile_handler.go   # Handler get & upsert profile
│
└── pkg/                     # Utility & helper packages
    ├── jwt.go               # Generate & Validate JWT token
    ├── password.go          # Hash & verifikasi password (bcrypt)
    ├── response.go          # Standarisasi format JSON response
    └── middleware/
        └── middleware.go    # JWT Auth middleware
```

---

## 🗄️ Database Schema (ERD)

```
┌──────────────┐
│    users     │
├──────────────┤
│ id (PK)      │──┐
│ email        │  │  1:1    ┌───────────────┐
│ password     │  ├────────▶│   profiles    │
│ created_at   │  │         ├───────────────┤
│ updated_at   │  │         │ user_id (FK)  │
│ deleted_at   │  │         │ full_name     │
└──────────────┘  │         │ phone         │
                  │         │ address       │
                  │         │ photo_url     │
                  │         │ summary       │
                  │         └───────────────┘
                  │
                  │  1:N    ┌────────────────────┐
                  ├────────▶│ work_experiences   │
                  │         ├────────────────────┤
                  │         │ company_name       │
                  │         │ position           │
                  │         │ start_date         │
                  │         │ end_date           │
                  │         │ is_current         │
                  │         │ description        │
                  │         └────────────────────┘
                  │
                  │  1:N    ┌──────────────────┐
                  ├────────▶│   educations     │
                  │         ├──────────────────┤
                  │         │ institution      │
                  │         │ degree           │
                  │         │ field_of_study   │
                  │         │ start_date       │
                  │         │ end_date         │
                  │         │ gpa              │
                  │         └──────────────────┘
                  │
                  │  1:N    ┌──────────────┐
                  ├────────▶│   skills     │
                  │         ├──────────────┤
                  │         │ name         │
                  │         │ level        │
                  │         └──────────────┘
                  │
                  │  1:N    ┌──────────────┐
                  └────────▶│  projects    │
                            ├──────────────┤
                            │ title        │
                            │ description  │
                            │ link         │
                            │ tech_stack   │
                            └──────────────┘
```

---

## 🚀 Cara Menjalankan

### Prasyarat

- Go >= 1.25
- PostgreSQL
- Git

### 1. Clone Repository

```bash
git clone https://github.com/rakafajar/cv_builder_be.git
cd cv_builder_be
```

### 2. Konfigurasi Environment

Buat file `.env` di root project:

```env
DB_HOST=localhost
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=cv_builder_db
DB_PORT=5432
JWT_SECRET=your_super_secret_key
```

### 3. Buat Database

```bash
createdb cv_builder_db
```

> Tabel akan di-buat otomatis saat aplikasi pertama kali dijalankan (auto-migrate oleh GORM).

### 4. Install Dependencies & Jalankan

```bash
go mod tidy
go run main.go
```

Server berjalan di `http://localhost:8080`

---

## 📡 API Endpoints

Base URL: `http://localhost:8080/api/v1`

### 🔓 Public (Tanpa Token)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| `POST` | `/register` | Registrasi user baru |
| `POST` | `/login` | Login & mendapatkan JWT token |

### 🔐 Protected (Wajib Token JWT)

Header: `Authorization: Bearer <token>`

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| `GET` | `/profile` | Mendapatkan data profile user |
| `POST` | `/profile` | Membuat atau mengupdate profile user |

---

## 📨 Contoh Request & Response

### Register

```bash
curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{"email": "john@mail.com", "password": "secret123"}'
```

**Response (201):**

```json
{
  "status": "Created",
  "response_code": 201,
  "message": "Registrasi Berhasil",
  "data": {
    "id": 1,
    "email": "john@mail.com"
  }
}
```

### Login

```bash
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"email": "john@mail.com", "password": "secret123"}'
```

**Response (200):**

```json
{
  "status": "Success",
  "response_code": 200,
  "message": "Login Berhasil",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

### Get Profile

```bash
curl http://localhost:8080/api/v1/profile \
  -H "Authorization: Bearer <token>"
```

**Response (200):**

```json
{
  "status": "Success",
  "response_code": 200,
  "message": "Berhasil Mendapatkan Profile",
  "data": {
    "user_id": 1,
    "full_name": "John Doe",
    "phone": "08123456789",
    "address": "Jakarta, Indonesia",
    "photo_url": "https://example.com/photo.jpg",
    "summary": "Backend Developer with 3 years of experience"
  }
}
```

### Upsert Profile (Create / Update)

```bash
curl -X POST http://localhost:8080/api/v1/profile \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "John Doe",
    "phone": "08123456789",
    "address": "Jakarta, Indonesia",
    "photo_url": "https://example.com/photo.jpg",
    "summary": "Backend Developer with 3 years of experience"
  }'
```

**Response (201):**

```json
{
  "status": "Created",
  "response_code": 201,
  "message": "Berhasil membuat data atau mengubah data profile"
}
```

---

## 🏗️ Arsitektur

Project ini menggunakan pola **Clean Architecture** yang memisahkan kode menjadi beberapa layer:

```
Request  →  [Delivery/Handler]  →  [Usecase]  →  [Repository]  →  Database
Response ←  [Delivery/Handler]  ←  [Usecase]  ←  [Repository]  ←  Database
```

| Layer | Tanggung Jawab |
|-------|----------------|
| **Delivery** | Menerima HTTP request, parsing input, memanggil usecase, menulis response |
| **Usecase** | Business logic (validasi, hashing password, generate token) |
| **Repository** | Query database melalui GORM |
| **Models** | Definisi struct yang merepresentasikan tabel database |
| **Pkg** | Utility bersama: JWT, password hashing, response helper, middleware |

---

## 📋 Format Response

Semua response API mengikuti format standar:

```json
{
  "status": "Success | Created | Error Bad Request | Error Unauthorized | Error Not Found",
  "response_code": 200,
  "message": "Deskripsi hasil",
  "data": {},
  "error": "Detail error (hanya jika ada error)"
}
```

---

## 🛡️ Autentikasi

1. User melakukan **Register** dengan email & password
2. Password di-hash menggunakan **bcrypt** sebelum disimpan
3. User melakukan **Login** dan mendapatkan **JWT Token** (berlaku 24 jam)
4. Token dikirim melalui header `Authorization: Bearer <token>` untuk mengakses endpoint protected
5. Middleware memvalidasi token dan meng-inject `userID` ke dalam request context

---

## 📝 License

MIT
