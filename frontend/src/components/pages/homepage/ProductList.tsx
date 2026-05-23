import ProductCard from './ProductCard';

interface IProduct {
  id: number;
  title: string;
  description: string;
  price: number;
  imageUrl: string;
}

export default async function ProductList() {
  const res = await fetch(`${process.env.API_URL}/products`, {
    cache: 'no-store',
  });
  const products = await res.json();

  return (
    <section className='p-8'>
      <h2 className='text-center text-3xl font-bold mb-4'>All Products</h2>
      <ul className='grid grid-cols-3 gap-10'>
        {products.map((p: IProduct) => (
          <ProductCard product={p} key={p.id} />
        ))}
      </ul>
    </section>
  );
}
