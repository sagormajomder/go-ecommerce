package database


type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

var productList []Product

func Store (p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product{
	return productList
}

func Get(id int) *Product{
	for _, product := range productList {
		 if product.ID == id {
			return &product
		 }
	}

	return nil
}

func Update(id int, p Product) Product {
	p.ID = id
	for idx, product := range productList {
		 if product.ID == p.ID {
			productList[idx] = p
		 }
	}

	return p
}

func Delete(id int){
	var tempList []Product
	for _, product := range productList {
		 if product.ID != id {
			tempList = append(tempList, product)
		 }
	}
	productList = tempList
}

func init() {
	productList = append(productList, Product{ID: 1, Title: "Wireless Headphones", Description: "High-quality noise-cancelling headphones.", Price: 129.99, ImgURL: "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS"},
		Product{ID: 2, Title: "Smart Watch", Description: "Stylish smart watch with health tracking.", Price: 199.99, ImgURL: "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528"},
		Product{ID: 3, Title: "Running Shoes", Description: "Lightweight shoes for everyday running.", Price: 89.50, ImgURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s"},
		Product{ID: 4, Title: "Wireless Headphones", Description: "High-quality noise-cancelling headphones.", Price: 129.99, ImgURL: "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS"},
		Product{ID: 5, Title: "Smart Watch", Description: "Stylish smart watch with health tracking.", Price: 199.99, ImgURL: "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528"},
		Product{ID: 6, Title: "Running Shoes", Description: "Lightweight shoes for everyday running.", Price: 89.50, ImgURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s"})
}