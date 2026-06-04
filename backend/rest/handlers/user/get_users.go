package user

import (
	"ecommerce/util"
	"fmt"
	"net/http"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {

	allUser, err := h.userRepo.List()

	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if allUser != nil {
		util.SendData(w, allUser, http.StatusOK)
		return
	}

	util.SendError(w, http.StatusNotFound, "No users created")
}
