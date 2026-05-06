# Product Management Frontend

![React](https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB)
![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white)
![Vite](https://img.shields.io/badge/vite-%23646CFF.svg?style=for-the-badge&logo=vite&logoColor=white)
![TailwindCSS](https://img.shields.io/badge/tailwindcss-%2338B2AC.svg?style=for-the-badge&logo=tailwind-css&logoColor=white)
![React Router](https://img.shields.io/badge/React_Router-CA4245?style=for-the-badge&logo=react-router&logoColor=white)

> Antarmuka web responsif untuk Sistem Manajemen Produk. Dibangun dengan fokus pada kebersihan antarmuka (*clean UI*), *readability* kode, dan arsitektur *Single Page Application* (SPA) yang terstruktur.

Aplikasi ini merupakan bagian *Frontend* yang terintegrasi langsung dengan *Backend* Go REST API untuk mengelola produk secara *real-time*.

---

## Alur Aplikasi

Aplikasi ini menerapkan siklus CRUD penuh tanpa perlu memuat ulang halaman:

1. **Katalog Produk - `/`**
   Saat aplikasi dibuka, sistem akan melakukan pengambilan data ke *Backend* dan menampilkan daftar produk dalam bentuk *card* yang responsif. Menampilkan status kosong jika belum ada data.
2. **Tambah Produk - `/add`**
   Menekan tombol `+ Tambah Produk` akan mengarahkan pengguna ke halaman form tambah produk. Setelah validasi dan penyimpanan sukses, pengguna otomatis dikembalikan ke halaman utama dengan daftar data terbaru.
3. **Edit Produk - `/edit/:id`**
   Setiap list produk memiliki tombol `Edit`. Sistem akan mengambil ID dari URL, menarik data produk tersebut dari *Backend*, dan mengisi form secara otomatis untuk diperbarui.
4. **Hapus Produk**
   Tindakan penghapusan dilengkapi dengan *dialog* konfirmasi untuk mencegah hilangnya data akibat ketidaksengajaan.

---

## Teknologi & Library

Aplikasi ini dikembangkan dengan *stack* modern untuk memastikan performa dan skalabilitas:

| Teknologi | Fungsi / Kegunaan |
| :--- | :--- |
| **React** | *Library* utama untuk merender komponen UI secara dinamis. |
| **TypeScript** | Memberikan *static typing* pada JavaScript. Mencegah *bug* sejak tahap penulisan kode dan memastikan struktur data `Product` sinkron dengan *Backend*. |
| **Vite** | Memberikan waktu *startup* local server yang sangat cepat dan proses *build* yang efisien. |
| **Tailwind CSS 3** | *Framework CSS* untuk menciptakan desain yang bersih, proporsional, dan sangat **responsif**. |
| **React Router** | Mengatur *routing* di sisi klien, memungkinkan perpindahan antalaman (`/`, `/add`, `/edit`) dengan mulus tanpa memuat ulang *browser*. |

---

## Struktur Direktori

Kode dipecah berdasarkan fungsionalitasnya agar mudah dipelihara:

```text
src/
├── api/
│   └── productApi.ts     # Sentralisasi panggilan HTTP (fetch) ke Backend Go.
├── components/
│   └── ProductForm.tsx   # Komponen formulir yang 'reusable' untuk Create dan Update.
├── pages/
│   ├── Home.tsx          # Halaman '/' -> Menampilkan daftar semua produk.
│   ├── AddProduct.tsx    # Halaman '/add' -> Halaman pembuatan produk baru.
│   └── EditProduct.tsx   # Halaman '/edit/:id' -> Halaman pengeditan produk.
├── types/
│   └── product.ts        # Definisi tipe data (Interface) standar TypeScript.
├── App.tsx               # Konfigurasi perutean (Router) utama aplikasi.
├── index.css             # Entry point untuk direktif Tailwind CSS.
└── main.tsx              # Entry point aplikasi React.
```

---

## Panduan Memulai Lokal (Local Setup)

Ikuti langkah berikut untuk menjalankan aplikasi ini di mesin Anda:

**1. Prasyarat**
Pastikan Anda telah menginstal [Node.js](https://nodejs.org/) dan *Backend* Go Anda sudah berjalan di `http://localhost:8080`.

**2. Instalasi & Menjalankan Server**
Jalankan perintah berikut secara berurutan di terminal Anda:

```bash
# Kloning repositori atau buka folder proyek
cd product-app-frontend

# Instal semua dependensi
npm install

# Jalankan server pengembangan (development server)
npm run dev
```

**3. Akses Aplikasi**
Buka *browser* Anda dan kunjungi `http://localhost:5173` (atau port yang tertera di terminal Anda).
