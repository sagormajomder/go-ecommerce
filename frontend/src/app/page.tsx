import ProductForm from '@/components/pages/homepage/ProductForm';
import ProductList from '@/components/pages/homepage/ProductList';

export default function Home() {
  return (
    <section className='py-14 space-y-10'>
      <ProductForm />
      <ProductList />
    </section>
  );
}
