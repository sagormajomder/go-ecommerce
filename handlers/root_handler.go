package handlers

import (
	"ecommerce/util"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request){
	util.SendData(w,"Hello World",200)
}