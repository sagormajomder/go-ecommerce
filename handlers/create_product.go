package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product

	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
		return
	}
	newProduct.ID = len(database.ProductList) + 1

	database.ProductList = append(database.ProductList, newProduct)

	util.SendData(w, newProduct, 201)

}