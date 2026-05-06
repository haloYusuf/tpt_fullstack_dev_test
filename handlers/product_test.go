package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"product-app-backend/database"
	"product-app-backend/repository"
	"product-app-backend/service"
)

// menguji apakah handler mengembalikan error 400 Bad Request
// jika ID yang diberikan di URL bukan berupa angka
func TestProductByIDHandler_InvalidID(t *testing.T) {
	// membuat handler dengan service 'nil' karena skenario test ini gagal dan berhenti di level Handler sebelum memanggil Service/Database.
	handler := NewProductHandler(nil)

	// Membuat request HTTP buatan dengan id "abc" yang tidak valid
	req, err := http.NewRequest("GET", "/products/abc", nil)
	if err != nil {
		t.Fatalf("Gagal membuat request: %v", err)
	}

	// Membuat ResponseRecorder untuk merekam balasan dari handler
	rr := httptest.NewRecorder()

	// Menjalankan handler dengan ResponseRecorder dan Request buatan
	handler.ProductByIDHandler(rr, req)

	// Memeriksa apakah status code yang dikembalikan adalah 400 Bad Request
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler mengembalikan status code yang salah: mendapat %v, seharusnya %v",
			status, http.StatusBadRequest)
	}

	// Memeriksa apakah pesan error di dalam body json sesuai dengan ekspektasi
	expected := `{"error":"ID produk harus berupa angka"}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("Handler mengembalikan body yang salah: mendapat %v, seharusnya %v",
			rr.Body.String(), expected)
	}
}

// Tmenguji apakah POST /products mengembalikan error 400 Bad Request jika format JSON yang dikirimkan rusak.
func TestProductsHandler_InvalidJSON(t *testing.T) {
	// Menggunakan nil untuk service karena request akan ditolak di level handler sebelum masuk ke service.
	handler := NewProductHandler(nil)

	// Membuat data JSON buatan yang sengaja dibuat cacat (tidak ada kurung kurawal penutup)
	badJSON := `{"name": "Kopi Susu", "price": 15000`
	body := strings.NewReader(badJSON)

	// Membuat request HTTP POST ke /products dengan body JSON yang rusak
	req, err := http.NewRequest("POST", "/products", body)
	if err != nil {
		t.Fatalf("Gagal membuat request: %v", err)
	}

	// Membuat ResponseRecorder untuk menangkap balasan
	rr := httptest.NewRecorder()

	// Eksekusi
	handler.ProductsHandler(rr, req)

	// status 400 Bad Request karena JSON-nya gagal di-decode
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler mengembalikan status code yang salah: mendapat %v, seharusnya %v",
			status, http.StatusBadRequest)
	}
}

// menguji skenario sukses (200 OK) saat mengambil semua daftar produk.
func TestProductsHandler_GetAll_Success(t *testing.T) {
	// Menyambungkannya ke database asli
	db := database.ConnectDB()
	// Tutup koneksi setelah tes selesai
	defer db.Close()

	repo := repository.NewProductRepository(db)
	svc := service.NewProductService(repo)
	handler := NewProductHandler(svc)

	// Membuat request HTTP GET ke /products
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatalf("Gagal membuat request: %v", err)
	}

	// Membuat ResponseRecorder untuk menangkap balasan
	rr := httptest.NewRecorder()

	// Eksekusi
	handler.ProductsHandler(rr, req)

	// Mengharapkan status 200 OK karena berhasil mengambil data
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler mengembalikan status code yang salah: mendapat %v, seharusnya %v",
			status, http.StatusOK)
	}
}
