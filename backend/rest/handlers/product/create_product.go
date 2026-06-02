package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct ReqCreateProduct

	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{
		Title:       newProduct.Title,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		ImgURL:      newProduct.ImgURL,
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, createdProduct, http.StatusCreated)

}
