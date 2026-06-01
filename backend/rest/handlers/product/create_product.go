package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
		return
	}

	createdProduct := database.Store(newProduct)

	util.SendData(w, createdProduct, 201)

}
