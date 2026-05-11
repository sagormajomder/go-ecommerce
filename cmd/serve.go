package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve(){
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductByID))

	port := ":8080"
	globalRouter := global_router.GlobalRouter(mux)
	println("🚀 Server is running at http://localhost" + port)

	err := http.ListenAndServe(port, globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}