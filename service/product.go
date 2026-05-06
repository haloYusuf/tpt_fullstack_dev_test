package service

import (
	"errors"
	"product-app-backend/models"
	"product-app-backend/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// Validasi data sebelum disimpan
func (s *ProductService) CreateProduct(p *models.Product) error {
	// Validasi Required Fields
	if p.Name == "" {
		return errors.New("nama produk tidak boleh kosong")
	}
	if p.Price <= 0 {
		return errors.New("harga produk harus lebih dari 0")
	}
	if p.Stock < 0 {
		return errors.New("stok tidak boleh negatif")
	}

	return s.repo.Create(p)
}

// GetAllProducts langsung memanggil repository
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAll()
}

// GetProductByID mengambil produk dengan validasi id
func (s *ProductService) GetProductByID(id int) (*models.Product, error) {
	if id <= 0 {
		return nil, errors.New("ID produk tidak valid")
	}
	return s.repo.GetByID(id)
}

// UpdateProduct memvalidasi data baru sebelum diubah
func (s *ProductService) UpdateProduct(id int, p *models.Product) error {
	if id <= 0 {
		return errors.New("ID produk tidak valid")
	}
	if p.Name == "" {
		return errors.New("nama produk tidak boleh kosong")
	}
	if p.Price <= 0 {
		return errors.New("harga produk harus lebih dari 0")
	}

	return s.repo.Update(id, p)
}

// DeleteProduct memvalidasi id sebelum menghapus
func (s *ProductService) DeleteProduct(id int) error {
	if id <= 0 {
		return errors.New("ID produk tidak valid")
	}
	return s.repo.Delete(id)
}
