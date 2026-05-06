package repository

import (
	"database/sql"
	"errors"
	"product-app-backend/models"
)

// Menaruh koneksi database untuk digunakan oleh fungsi-fungsi di bawahnya
type ProductRepository struct {
	db *sql.DB
}

// Fungsi untuk membuat pemanggilan repository baru
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Menambahkan produk baru ke database
func (r *ProductRepository) Create(p *models.Product) error {
	query := `
		INSERT INTO products (name, description, price, stock, category) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id, is_active, created_at, updated_at`

	// QueryRow digunakan karena ingin return data
	err := r.db.QueryRow(query, p.Name, p.Description, p.Price, p.Stock, p.Category).
		Scan(&p.ID, &p.IsActive, &p.CreatedAt, &p.UpdatedAt)
	return err
}

// Mengambil semua produk
func (r *ProductRepository) GetAll() ([]models.Product, error) {
	query := `SELECT id, name, description, price, stock, category, is_active, created_at, updated_at FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.Category, &p.IsActive, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// Mengambil satu produk berdasarkan ID
func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
	query := `SELECT id, name, description, price, stock, category, is_active, created_at, updated_at FROM products WHERE id = $1`
	var p models.Product

	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.Category, &p.IsActive, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("produk tidak ditemukan")
		}
		return nil, err
	}
	return &p, nil
}

// Mengubah data produk
func (r *ProductRepository) Update(id int, p *models.Product) error {
	query := `
		UPDATE products 
		SET name = $1, description = $2, price = $3, stock = $4, category = $5, updated_at = CURRENT_TIMESTAMP
		WHERE id = $6`

	result, err := r.db.Exec(query, p.Name, p.Description, p.Price, p.Stock, p.Category, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New("produk tidak ditemukan atau tidak ada perubahan")
	}
	return nil
}

// Menghapus produk
func (r *ProductRepository) Delete(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New("produk tidak ditemukan")
	}
	return nil
}
