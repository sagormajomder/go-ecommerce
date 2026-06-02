package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	var updateProduct ReqUpdateProduct

	err = json.NewDecoder(r.Body).Decode(&updateProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
		return
	}

	_, err = h.productRepo.Update(id, repo.Product{
		Title:       updateProduct.Title,
		Description: updateProduct.Description,
		Price:       updateProduct.Price,
		ImgURL:      updateProduct.ImgURL,
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, "Product Updated Successfully", 201)

}
