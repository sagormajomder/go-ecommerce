package user

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {

	allUser, err := h.userRepo.List()

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if allUser != nil {
		util.SendData(w, allUser, http.StatusOK)
		return
	}

	http.Error(w, "No users created", http.StatusNotFound)
}
