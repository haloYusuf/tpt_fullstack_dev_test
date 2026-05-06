package database

import (
	"database/sql"
	"fmt"
	"log"

	// Import driver postgres
	_ "github.com/lib/pq"
)

// ConnectDB membuka koneksi ke database dan membuat tabel (Jika belum ada)
func ConnectDB() *sql.DB {
	// Kredensial sesuai file compose.yml
	connStr := "postgres://root:secretpassword@localhost:5432/product_db?sslmode=disable"

	db, error := sql.Open("postgres", connStr)
	if error != nil {
		log.Fatalf("Gagal membuka database: %v", error)
	}

	// Cek koneksi ke database
	error = db.Ping()
	if error != nil {
		log.Fatalf("Gagal terhubung ke database: %v", error)
	}

	fmt.Println("Berhasil terhubung ke PostgreSQL!")
	createTable(db)

	return db
}

// createTable menyiapkan skema database otomatis
func createTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		price NUMERIC(10, 2) NOT NULL,
		stock INT NOT NULL,
		category VARCHAR(100),
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Gagal membuat tabel: %v", err)
	}
	fmt.Println("Tabel 'products' siap digunakan.")
}
