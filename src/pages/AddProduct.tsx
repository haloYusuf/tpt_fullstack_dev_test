import { useNavigate } from 'react-router-dom';
import { productApi } from '../api/productApi';
import ProductForm from '../components/ProductForm';
import type { Product } from '../types/product';

export default function AddProduct() {
    const navigate = useNavigate();

    const handleCreate = async (productData: Product) => {
        try {
            await productApi.create(productData);
            alert('Produk berhasil ditambahkan!');
            // Kembali ke Home
            navigate('/');
        } catch (error) {
            alert('Gagal menambahkan produk.');
        }
    };

    return (
        <div className="min-h-screen bg-gray-50 p-3 sm:p-8">
            <div className="max-w-2xl mx-auto">
                <ProductForm
                    onSubmit={handleCreate}
                    // Kembali ke Home
                    onCancel={() => navigate('/')}
                />
            </div>
        </div>
    );
}