package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	var updateProduct database.Product

	err = json.NewDecoder(r.Body).Decode(&updateProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
		return
	}

	database.Update(id, updateProduct)

	util.SendData(w, "Product Updated Successfully", 201)

}
