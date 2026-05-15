package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve(){
	mux := http.NewServeMux()


	// mux.Handle("GET /", middleware.Logger(http.HandlerFunc(handlers.GetRoot))) 
	manager:= middleware.NewManager()
	logger:=manager.With(middleware.Logger)
	mux.Handle("GET /", logger(http.HandlerFunc(handlers.GetRoot))) 


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