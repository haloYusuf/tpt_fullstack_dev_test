package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"product-app-backend/models"
	"product-app-backend/service"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// helper untuk respon dengan format json
func responseJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// helper untuk respon pesan error
func responseError(w http.ResponseWriter, status int, message string) {
	responseJSON(w, status, map[string]string{"error": message})
}

// Handler menangani rute GET /products dan POST /products tanpa param
func (h *ProductHandler) ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		products, err := h.service.GetAllProducts()
		if err != nil {
			responseError(w, http.StatusInternalServerError, err.Error())
			return
		}
		responseJSON(w, http.StatusOK, products)

	case http.MethodPost:
		var p models.Product
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			responseError(w, http.StatusBadRequest, "Format JSON tidak valid")
			return
		}

		if err := h.service.CreateProduct(&p); err != nil {
			responseError(w, http.StatusBadRequest, err.Error())
			return
		}
		responseJSON(w, http.StatusCreated, p)

	default:
		responseError(w, http.StatusMethodNotAllowed, "Metode tidak diizinkan")
	}
}

// Handler menangani rute dengan param / id (GET, PUT, DELETE /products/:id)
func (h *ProductHandler) ProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil id dari url
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[2] == "" {
		responseError(w, http.StatusBadRequest, "ID produk wajib disertakan")
		return
	}

	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		responseError(w, http.StatusBadRequest, "ID produk harus berupa angka")
		return
	}

	switch r.Method {
	case http.MethodGet:
		product, err := h.service.GetProductByID(id)
		if err != nil {
			// Menangani error not found
			if err.Error() == "produk tidak ditemukan" {
				responseError(w, http.StatusNotFound, err.Error()) // 404
				return
			}
			responseError(w, http.StatusInternalServerError, err.Error())
			return
		}
		responseJSON(w, http.StatusOK, product)

	case http.MethodPut:
		var p models.Product
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			responseError(w, http.StatusBadRequest, "Format JSON tidak valid")
			return
		}

		if err := h.service.UpdateProduct(id, &p); err != nil {
			if err.Error() == "produk tidak ditemukan atau tidak ada perubahan" {
				responseError(w, http.StatusNotFound, err.Error())
				return
			}
			responseError(w, http.StatusBadRequest, err.Error())
			return
		}
		responseJSON(w, http.StatusOK, map[string]string{"message": "Produk berhasil diupdate"})

	case http.MethodDelete:
		if err := h.service.DeleteProduct(id); err != nil {
			if err.Error() == "produk tidak ditemukan" {
				responseError(w, http.StatusNotFound, err.Error())
				return
			}
			responseError(w, http.StatusInternalServerError, err.Error())
			return
		}
		responseJSON(w, http.StatusOK, map[string]string{"message": "Produk berhasil dihapus"})

	default:
		responseError(w, http.StatusMethodNotAllowed, "Metode tidak diizinkan")
	}
}
