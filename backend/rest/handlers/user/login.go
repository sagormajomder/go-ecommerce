package user

import (
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var reqLogin ReqLogin

	err := json.NewDecoder(r.Body).Decode(&reqLogin)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	usr, err := h.userRepo.Get(reqLogin.Email, reqLogin.Password)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if usr == nil {
		http.Error(w, "Invalid Credential", http.StatusBadRequest)
		return
	}

	accessToken, err := util.CreateJWT(h.cnf.JWTSecret, util.Payload{
		Sub:         usr.ID,
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	})

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, accessToken, http.StatusOK)

}
