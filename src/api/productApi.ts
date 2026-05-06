import type { Product } from "../types/product";

const BASE_URL = "http://localhost:8080/products";

export const productApi = {
  // Mengambil semua produk
  async getAll(): Promise<Product[]> {
    const response = await fetch(BASE_URL);
    if (!response.ok) throw new Error("Gagal mengambil data");
    return response.json();
  },

  // Mengambil produk berdasarkan id
  async getById(id: number): Promise<Product> {
    const response = await fetch(`${BASE_URL}/${id}`);
    if (!response.ok) throw new Error("Produk tidak ditemukan");
    return response.json();
  },

  // Membuat produk baru
  async create(product: Product): Promise<Product> {
    const response = await fetch(BASE_URL, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(product),
    });
    if (!response.ok) throw new Error("Gagal menambah produk");
    return response.json();
  },

  // Mengupdate produk
  async update(id: number, product: Product): Promise<void> {
    const response = await fetch(`${BASE_URL}/${id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(product),
    });
    if (!response.ok) throw new Error("Gagal mengupdate produk");
  },

  // Menghapus produk
  async delete(id: number): Promise<void> {
    const response = await fetch(`${BASE_URL}/${id}`, {
      method: "DELETE",
    });
    if (!response.ok) throw new Error("Gagal menghapus produk");
  },
};
