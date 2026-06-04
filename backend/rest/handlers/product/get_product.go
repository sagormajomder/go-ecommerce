package product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)

	if err != nil {
		fmt.Println(err)
		util.SendError(w, 400, "Please give me a valid product id")
		return
	}

	productById, err := h.productRepo.Get(id)

	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if productById == nil {
		util.SendError(w, 404, "Product not found")
		return
	}

	util.SendData(w, productById, 200)

}
