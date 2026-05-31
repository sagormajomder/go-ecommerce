package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser database.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	createdUser := newUser.Store()

	if createdUser == nil {
		http.Error(w, "User Already Existed", http.StatusBadRequest)
		return
	}

	util.SendData(w, createdUser, http.StatusCreated)

}
