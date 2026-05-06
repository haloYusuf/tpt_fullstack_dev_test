# Product Management App - Backend

Dokumentasi ini dibuat untuk memenuhi tugas seleksi Fullstack Developer Intern. Backend ini dibangun menggunakan bahasa **Go** tanpa framework eksternal (menggunakan `net/http`) dan menggunakan **PostgreSQL** sebagai database utama.

## Fitur Utama

- RESTful API untuk pengelolaan produk (CRUD).
- Validasi field wajib (Name, Price, Stock).
- Penanganan error HTTP status code.
- Unit Testing.
- Database PostgreSQL yang berjalan di dalam Docker.

---

## Struktur Folder

Proyek ini menggunakan pola pemisahan tanggung jawab agar mudah dipahami:

- **`main.go`**: Entry Point aplikasi. Menangani inisialisasi koneksi, routing, dan menjalankan server.
- **`database/`**: Berisi konfigurasi koneksi ke PostgreSQL dan skema pembuatan tabel otomatis.
- **`handlers/`**: Menangani Request dan Response HTTP. Membaca JSON dan mengirimkan data kembali.
- **`models/`**: Blueprint Produk yang digunakan di seluruh aplikasi.
- **`repository/`**: Lapisan yang berhubungan langsung dengan database (menjalankan query SQL).
- **`service/`**: Lapisan logika bisnis. Tempat dilakukan validasi data sebelum dikirim ke repository.

---

## Prasyarat

Sebelum menjalankan aplikasi, pastikan Anda sudah menginstal:

- [Go](https://go.dev/dl/) (Versi 1.21 atau lebih baru)
- [Docker](https://www.docker.com/get-started) & [Docker Compose](https://docs.docker.com/compose/install/)

---

## Cara Menjalankan Aplikasi

### 1. Jalankan Database (Docker)

Pastikan Docker Anda menyala, lalu jalankan container PostgreSQL:

```bash
docker compose up -d
```

### 2. Persiapkan Dependency Go

Unduh driver database yang diperlukan:

```bash
go mod tidy
```

### 3. Jalankan Backend

Unduh driver database yang diperlukan:

```bash
go mod tidy
```

**3. Jalankan Server Backend**
Jalankan perintah ini untuk menyalakan API:

```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`.

## Cara Menjalankan Unit Test

Proyek ini menyertakan _unit test_ menggunakan `net/http/httptest` untuk menguji validasi dan respons HTTP di level handler. Jalankan perintah berikut di terminal:

```bash
go test -v ./handlers
```

## Dokumentasi API Sederhana

**Base URL:** `http://localhost:8080`

### 1. Mengambil Semua Produk

- **Method:** `GET`
- **Endpoint:** `/products`
- **Response (200 OK):**
  ```json
  [
    {
      "id": 1,
      "name": "Kopi",
      "description": "Kopi susu",
      "price": 15000,
      "stock": 50,
      "category": "Minuman",
      "isActive": true,
      "createdAt": "2023-10-25T10:00:00Z",
      "updatedAt": "2023-10-25T10:00:00Z"
    }
  ]
  ```

### 2. Mengambil Produk Berdasarkan ID

- **Method:** `GET`
- **Endpoint:** `/products/:id`
- **Response (200 OK):** Objek JSON.
- **Response (404 Not Found):** `{"error": "produk tidak ditemukan"}`

### 3. Menambahkan Produk Baru

- **Method:** `POST`
- **Endpoint:** `/products`
- **Request Body:**
  ```json
  {
    "name": "Kopi",
    "description": "Kopi susu",
    "price": 15000,
    "stock": 50,
    "category": "Minuman"
  }
  ```
- **Response (201 Created):** Objek JSON.
- **Response (400 Bad Request):** Jika field wajib tidak diisi, format JSON salah, atau tipe data tidak sesuai.

### 4. Mengubah Data Produk

- **Method:** `PUT`
- **Endpoint:** `/products/:id`
- **Request Body (JSON):**
  ```json
  {
    "name": "Kopi",
    "description": "Kopi susu",
    "price": 15000,
    "stock": 49,
    "category": "Minuman"
  }
  ```
- **Response (200 OK):** `{"message": "Produk berhasil diupdate"}`
- **Response (404 Not Found):** Jika ID produk tidak ditemukan di database.

### 5. Menghapus Produk

- **Method:** `DELETE`
- **Endpoint:** `/products/:id`
- **Response (200 OK):** `{"message": "Produk berhasil dihapus"}`
- **Response (404 Not Found):** Jika ID produk tidak ditemukan di database.
