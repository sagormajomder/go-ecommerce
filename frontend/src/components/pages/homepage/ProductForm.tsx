export default function ProductForm() {
  return (
    <section className='p-8 border border-amber-100 rounded-md'>
      <h1 className='text-center text-4xl mb-10'>Add a Product</h1>
      <form action='' className='flex flex-col gap-8'>
        <div className='grid grid-cols-2'>
          <label htmlFor='product-name'>Product Name: </label>
          <input
            className='px-2 py-3 bg-amber-100 focus-within:bg-amber-200 focus-within:outline-none rounded-md text-black'
            type='text'
            name='product-name'
            id='product-name'
          />
        </div>
        <div className='grid grid-cols-2'>
          <label htmlFor='product-desc'>Product Description: </label>
          <textarea
            className='px-2 py-3 bg-amber-100 focus-within:bg-amber-200 focus-within:outline-none rounded-md text-black'
            name='product-desc'
            id='product-desc'></textarea>
        </div>
        <div className='grid grid-cols-2'>
          <label htmlFor='product-price'>Product Price: </label>
          <input
            className='px-2 py-3 bg-amber-100 focus-within:bg-amber-200 focus-within:outline-none rounded-md text-black'
            type='number'
            name='product-price'
            id='product-price'
          />
        </div>
        <div className='grid grid-cols-2'>
          <label htmlFor='product-img-url'>Product Image url: </label>
          <input
            className='px-2 py-3 bg-amber-100 focus-within:bg-amber-200 focus-within:outline-none rounded-md text-black'
            type='url'
            name='product-img-url'
            id='product-img-url'
          />
        </div>
        <button
          className='px-6 py-3 bg-amber-50 text-black font-bold rounded-md'
          type='submit'>
          Submit
        </button>
      </form>
    </section>
  );
}
