package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request){

	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)

	if err !=nil{
		http.Error(w, "Please give me a valid product id",400)
		return
	}


	database.Delete(id)

	util.SendData(w, "Product Deleted Successfully",200)

}