package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgURL      string  `json:"imageUrl" db:"image_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Update(id int, p Product) (*Product, error)
	Delete(id int) error
}

type productRepo struct {
	// productList []*Product
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
	// generateInitialProducts(repo)
	// return repo
}

func (repo *productRepo) Create(p Product) (*Product, error) {

	query := `
	INSERT INTO products(
		title, 
		description, 
		price, 
		image_url
	) VALUES(
		$1,
		$2,
		$3,
		$4
	)
	RETURNING id
	`

	row := repo.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgURL)
	err := row.Scan(&p.ID)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo *productRepo) Get(id int) (*Product, error) {

	var prd Product

	query := `
	SELECT 
	id,
	title,
	description,
	price, 
	image_url 
	FROM products 
	WHERE id=$1;
	`
	err := repo.db.Get(&prd, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &prd, nil
}

func (repo *productRepo) List() ([]*Product, error) {
	var prdList []*Product

	query := `
	SELECT 
	id,
	title,
	description,
	price, 
	image_url 
	FROM products;
	`
	err := repo.db.Select(&prdList, query)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return prdList, nil
}

func (repo *productRepo) Delete(id int) error {
	query := `
	DELETE FROM products 
	WHERE id=$1
	`

	_, err := repo.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (repo *productRepo) Update(id int, p Product) (*Product, error) {

	query := `
	UPDATE products SET 
	title=$1,
	description=$2,
	price=$3,
	image_url=$4
	WHERE id=$5
	RETURNING id
	`
	row := repo.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgURL, id)
	err := row.Err()

	if err != nil {
		return nil, err
	}

	return &p, nil
}

// func generateInitialProducts(r *productRepo) {
// 	r.productList = append(r.productList, &Product{ID: 1, Title: "Wireless Headphones", Description: "High-quality noise-cancelling headphones.", Price: 129.99, ImgURL: "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS"},
// 		&Product{ID: 2, Title: "Smart Watch", Description: "Stylish smart watch with health tracking.", Price: 199.99, ImgURL: "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528"},
// 		&Product{ID: 3, Title: "Running Shoes", Description: "Lightweight shoes for everyday running.", Price: 89.50, ImgURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s"},
// 		&Product{ID: 4, Title: "Wireless Headphones", Description: "High-quality noise-cancelling headphones.", Price: 129.99, ImgURL: "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS"},
// 		&Product{ID: 5, Title: "Smart Watch", Description: "Stylish smart watch with health tracking.", Price: 199.99, ImgURL: "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528"},
// 		&Product{ID: 6, Title: "Running Shoes", Description: "Lightweight shoes for everyday running.", Price: 89.50, ImgURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s"})
// }
