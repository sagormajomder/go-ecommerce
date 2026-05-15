package middleware

import (
	"log"
	"net/http"
)


func ArektaHudai(next http.Handler) http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println("Ami arekta middleware")

		next.ServeHTTP(w,r)
	})

}