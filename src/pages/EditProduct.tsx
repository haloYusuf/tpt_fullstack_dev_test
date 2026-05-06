import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { productApi } from '../api/productApi';
import ProductForm from '../components/ProductForm';
import type { Product } from '../types/product';

export default function EditProduct() {
    const navigate = useNavigate();
    // Mengambil ID dari URL (/edit/:id)
    const { id } = useParams();

    const [initialData, setInitialData] = useState<Product | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchProduct = async () => {
            try {
                if (id) {
                    const data = await productApi.getById(Number(id));
                    setInitialData(data);
                }
            } catch (error) {
                alert('Gagal mengambil data produk');
                navigate('/');
            } finally {
                setLoading(false);
            }
        };
        fetchProduct();
    }, [id, navigate]);

    const handleUpdate = async (productData: Product) => {
        try {
            if (id) {
                await productApi.update(Number(id), productData);
                alert('Produk berhasil diperbarui!');
                // Kembali ke Home
                navigate('/');
            }
        } catch (error) {
            alert('Gagal memperbarui produk.');
        }
    };

    if (loading) {
        return <div className="p-8 text-center text-gray-500">Memuat data...</div>;
    }

    return (
        <div className="min-h-screen bg-gray-50 p-3 sm:p-8">
            <div className="max-w-2xl mx-auto">
                <ProductForm
                    initialData={initialData}
                    onSubmit={handleUpdate}
                    // Kembali ke Home
                    onCancel={() => navigate('/')}
                />
            </div>
        </div>
    );
}