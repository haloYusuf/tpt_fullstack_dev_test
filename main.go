package main

import (
	"fmt"
	"log"
	"net/http"

	"product-app-backend/database"
	"product-app-backend/handlers"
	"product-app-backend/repository"
	"product-app-backend/service"
)

// Middleware untuk mengizinkan CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mengizinkan akses (jika deploy nanti, ganti "*" dengan url Frontend)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Jika request adalah OPTIONS, langsung balas OK
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Inisialisasi Database
	db := database.ConnectDB()
	// Database ditutup saat aplikasi berhenti
	defer db.Close()

	// Dependency Injection
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Konfigurasi Router
	mux := http.NewServeMux()

	// Mendaftarkan Endpoint
	// /products untuk GET (All) dan POST
	mux.HandleFunc("/products", productHandler.ProductsHandler)

	// /products/ untuk GET by ID, PUT, dan DELETE
	mux.HandleFunc("/products/", productHandler.ProductByIDHandler)

	// Menyalakan Server
	port := ":8080"
	fmt.Printf("Server berjalan di http://localhost%s\n", port)

	if err := http.ListenAndServe(port, enableCORS(mux)); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
