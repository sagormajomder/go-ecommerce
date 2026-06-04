package product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)

	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Please give me a valid product id")
		return
	}

	err = h.productRepo.Delete(id)

	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
	}

	util.SendData(w, "Product Deleted Successfully", 200)

}
