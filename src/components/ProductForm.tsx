import React, { useState, useEffect } from 'react';
import type { Product } from '../types/product';

interface ProductFormProps {
    initialData?: Product | null;
    onSubmit: (product: Product) => Promise<void>;
    onCancel: () => void;
}

export default function ProductForm({ initialData, onSubmit, onCancel }: ProductFormProps) {
    const [formData, setFormData] = useState<Product>({
        name: '',
        description: '',
        price: 0,
        stock: 0,
        category: ''
    });

    const [loading, setLoading] = useState(false);

    useEffect(() => {
        if (initialData) {
            setFormData(initialData);
        }
    }, [initialData]);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value } = e.target;
        setFormData({
            ...formData,
            [name]: name === 'price' || name === 'stock' ? Number(value) : value
        });
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setLoading(true);
        try {
            await onSubmit(formData);
        } catch (error) {
            alert('Terjadi kesalahan saat menyimpan data');
        } finally {
            setLoading(false);
        }
    };

    return (
        <form onSubmit={handleSubmit} className="bg-white p-4 sm:p-6 rounded-lg shadow-md border border-gray-100">
            <h2 className="text-xl sm:text-2xl font-bold mb-4 sm:mb-6 text-gray-800">
                {initialData ? 'Edit Produk' : 'Tambah Produk Baru'}
            </h2>

            <div className="mb-4">
                <label className="block text-gray-700 text-sm sm:text-base font-medium mb-1.5">Nama Produk</label>
                <input required type="text" name="name" value={formData.name} onChange={handleChange}
                    className="w-full border border-gray-300 p-2 sm:p-2.5 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 transition" />
            </div>

            <div className="mb-4">
                <label className="block text-gray-700 text-sm sm:text-base font-medium mb-1.5">Deskripsi</label>
                <textarea name="description" value={formData.description} onChange={handleChange} rows={3}
                    className="w-full border border-gray-300 p-2 sm:p-2.5 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 transition" />
            </div>

            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-4">
                <div>
                    <label className="block text-gray-700 text-sm sm:text-base font-medium mb-1.5">Harga (Rp)</label>
                    <input required type="number" name="price" value={formData.price} onChange={handleChange} min="0"
                        className="w-full border border-gray-300 p-2 sm:p-2.5 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 transition" />
                </div>
                <div>
                    <label className="block text-gray-700 text-sm sm:text-base font-medium mb-1.5">Stok</label>
                    <input required type="number" name="stock" value={formData.stock} onChange={handleChange} min="0"
                        className="w-full border border-gray-300 p-2 sm:p-2.5 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 transition" />
                </div>
            </div>

            <div className="mb-6">
                <label className="block text-gray-700 text-sm sm:text-base font-medium mb-1.5">Kategori</label>
                <input required type="text" name="category" value={formData.category} onChange={handleChange}
                    className="w-full border border-gray-300 p-2 sm:p-2.5 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 transition" />
            </div>

            <div className="flex flex-col-reverse sm:flex-row justify-end gap-2 sm:gap-3 mt-6">
                <button type="button" onClick={onCancel}
                    className="w-full sm:w-auto px-4 py-2 bg-gray-100 text-gray-700 font-medium rounded-md hover:bg-gray-200 transition border border-gray-300">
                    Batal
                </button>
                <button type="submit" disabled={loading}
                    className="w-full sm:w-auto px-4 py-2 bg-blue-600 text-white font-medium rounded-md hover:bg-blue-700 transition disabled:opacity-50">
                    {loading ? 'Menyimpan...' : 'Simpan Produk'}
                </button>
            </div>
        </form>
    );
}