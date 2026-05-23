import Image from 'next/image';

interface IProduct {
  id: number;
  title: string;
  description: string;
  price: number;
  imageUrl: string;
}

export default function ProductCard({ product }: { product: IProduct }) {
  const { title, price, imageUrl, description } = product;

  return (
    <li className='rounded-lg overflow-hidden bg-white shadow-md text-black'>
      <div className='w-full h-48 relative'>
        <Image
          src={imageUrl}
          alt={title}
          fill
          sizes='(max-width: 768px) 100vw, 33vw'
          className='object-cover'
          loading='eager'
        />
      </div>
      <div className='p-4 flex flex-col gap-3'>
        <h3 className='text-xl font-semibold'>{title}</h3>
        <p className='text-sm text-gray-600 line-clamp-3'>{description}</p>
        <div className='mt-2 flex items-center justify-between'>
          <span className='font-bold text-lg'>TK. {price} BDT</span>
          <button className='px-3 py-1 bg-amber-500 text-white rounded-md text-sm'>
            Buy
          </button>
        </div>
      </div>
    </li>
  );
}
