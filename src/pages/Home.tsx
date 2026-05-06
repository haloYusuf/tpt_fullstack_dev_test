import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { productApi } from '../api/productApi';
import type { Product } from '../types/product';

export default function Home() {
    const [products, setProducts] = useState<Product[]>([]);
    const [loading, setLoading] = useState(true);
    // Fungsi navigasi
    const navigate = useNavigate();

    useEffect(() => {
        loadProducts();
    }, []);

    const loadProducts = async () => {
        setLoading(true);
        try {
            const data = await productApi.getAll();
            setProducts(Array.isArray(data) ? data : []);
        } catch (error) {
            console.error("Gagal mengambil data:", error);
            setProducts([]);
        } finally {
            setLoading(false);
        }
    };

    const handleDeleteProduct = async (id: number | undefined) => {
        if (!id) return;
        if (!window.confirm("Apakah Anda yakin ingin menghapus produk ini?")) return;

        try {
            await productApi.delete(id);
            alert('Produk berhasil dihapus!');
            loadProducts();
        } catch (error) {
            alert('Gagal menghapus produk.');
        }
    };

    return (
        <div className="min-h-screen bg-gray-50 p-3 sm:p-8">
            <div className="max-w-4xl mx-auto">
                <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-3 sm:gap-0 mb-5 sm:mb-8">
                    <div>
                        <h1 className="text-xl sm:text-3xl font-bold text-gray-800 leading-tight">Manajemen Produk</h1>
                        <p className="text-xs sm:text-sm text-gray-500 mt-1">Kelola daftar inventaris produk Anda</p>
                    </div>
                    <button
                        onClick={() => navigate('/add')}
                        className="w-full sm:w-auto bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 sm:px-5 sm:py-2.5 rounded-md text-sm sm:text-base font-medium transition shadow-sm"
                    >
                        + Tambah Produk
                    </button>
                </div>

                {loading ? (
                    <div className="flex justify-center items-center py-20">
                        <div className="animate-pulse text-gray-500 text-sm sm:text-base font-medium">Memuat data produk...</div>
                    </div>
                ) : products.length === 0 ? (
                    <div className="text-center bg-white p-8 sm:p-10 rounded-xl shadow-sm border border-gray-100 text-gray-500">
                        <p className="text-base sm:text-lg font-medium text-gray-600 mb-2">Inventaris Kosong</p>
                        <p className="text-xs sm:text-sm">Belum ada produk yang ditambahkan. Silakan klik tombol "Tambah Produk".</p>
                    </div>
                ) : (
                    <div className="grid gap-3">
                        {products.map((p) => (
                            <div key={p.id} className="bg-white p-3.5 sm:p-5 rounded-xl shadow-sm border border-gray-100 flex flex-col sm:flex-row justify-between items-start sm:items-center gap-3 sm:gap-4 hover:shadow-md transition duration-200">
                                <div className="w-full">
                                    <h2 className="text-base sm:text-xl font-bold text-gray-800 leading-snug line-clamp-1">{p.name}</h2>
                                    <p className="text-gray-500 text-xs sm:text-sm mb-2.5 mt-0.5 line-clamp-2">{p.description}</p>

                                    <div className="flex flex-wrap gap-1.5 sm:gap-2 text-[11px] sm:text-sm">
                                        <span className="bg-gray-100 text-gray-700 px-2 py-0.5 sm:py-1 rounded-full font-medium border border-gray-200">{p.category}</span>
                                        <span className="bg-green-50 text-green-700 px-2 py-0.5 sm:py-1 rounded-full font-bold border border-green-200">Rp {p.price?.toLocaleString('id-ID')}</span>
                                        <span className={`px-2 py-0.5 sm:py-1 rounded-full font-medium border ${p.stock < 10 ? 'bg-red-50 text-red-700 border-red-200' : 'bg-blue-50 text-blue-700 border-blue-200'}`}>Stok: {p.stock}</span>
                                    </div>
                                </div>

                                <div className="flex gap-2 w-full sm:w-auto mt-1 pt-3 sm:mt-0 sm:pt-0 border-t sm:border-0 border-gray-100 shrink-0">
                                    <button
                                        onClick={() => navigate(`/edit/${p.id}`)}
                                        className="flex-1 sm:flex-none text-center text-blue-600 hover:text-blue-800 bg-blue-50 hover:bg-blue-100 px-3 py-1.5 sm:px-4 sm:py-2 rounded-md transition text-xs sm:text-sm font-medium"
                                    >
                                        Edit
                                    </button>
                                    <button
                                        onClick={() => handleDeleteProduct(p.id)}
                                        className="flex-1 sm:flex-none text-center text-red-600 hover:text-red-800 bg-red-50 hover:bg-red-100 px-3 py-1.5 sm:px-4 sm:py-2 rounded-md transition text-xs sm:text-sm font-medium"
                                    >
                                        Hapus
                                    </button>
                                </div>
                            </div>
                        ))}
                    </div>
                )}
            </div>
        </div>
    );
}