package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	allUser := database.GetUsers()

	util.SendData(w, allUser, http.StatusOK)

}
