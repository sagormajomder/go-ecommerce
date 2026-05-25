package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	productById := database.Get(id)

	if productById == nil {
		util.SendError(w, 404, "Product not found")
		return
	}

	util.SendData(w, productById, 200)

}
