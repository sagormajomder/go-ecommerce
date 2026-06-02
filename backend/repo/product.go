package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Update(id int, p Product) (*Product, error)
	Delete(id int) error
}

type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func (repo *productRepo) Create(p Product) (*Product, error) {

	p.ID = len(repo.productList) + 1
	repo.productList = append(repo.productList, &p)
	return &p, nil
}

func (repo *productRepo) Get(id int) (*Product, error) {
	for _, product := range repo.productList {
		if product.ID == id {
			return product, nil
		}
	}

	return nil, nil
}
func (repo *productRepo) List() ([]*Product, error) {
	return repo.productList, nil
}

func (repo *productRepo) Delete(id int) error {
	var tempList []*Product
	for _, product := range repo.productList {
		if product.ID != id {
			tempList = append(tempList, product)
		}
	}
	repo.productList = tempList

	return nil
}

func (repo *productRepo) Update(id int, p Product) (*Product, error) {

	p.ID = id
	for idx, product := range repo.productList {
		if product.ID == p.ID {
			repo.productList[idx] = &p
		}
	}

	return &p, nil
}

func generateInitialProducts(r *productRepo) {
	r.productList = append(r.productList, &Product{ID: 1, Title: "Wireless Headphones", Description: "High-quality noise-cancelling headphones.", Price: 129.99, ImgURL: "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS"},
		&Product{ID: 2, Title: "Smart Watch", Description: "Stylish smart watch with health tracking.", Price: 199.99, ImgURL: "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528"},
		&Product{ID: 3, Title: "Running Shoes", Description: "Lightweight shoes for everyday running.", Price: 89.50, ImgURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s"},
		&Product{ID: 4, Title: "Wireless Headphones", Description: "High-quality noise-cancelling headphones.", Price: 129.99, ImgURL: "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS"},
		&Product{ID: 5, Title: "Smart Watch", Description: "Stylish smart watch with health tracking.", Price: 199.99, ImgURL: "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528"},
		&Product{ID: 6, Title: "Running Shoes", Description: "Lightweight shoes for everyday running.", Price: 89.50, ImgURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s"})
}
