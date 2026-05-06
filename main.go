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

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
