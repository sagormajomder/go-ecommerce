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


	// mux.Handle("GET /", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.GetRoot)))) 
	//* First Manager Implement 
	// manager:= middleware.NewManager()
	// mux.Handle("GET /", manager.With(
	// 										middleware.Hudai,
	// 										middleware.Logger,
	// 										)(http.HandlerFunc(handlers.GetRoot))) 

	//* Second Manager Implement
	// manager:= middleware.NewManager()
	// mux.Handle("GET /", manager.With(
	// 										http.HandlerFunc(handlers.GetRoot),
	// 										middleware.Logger, middleware.Hudai))

	//* Best Manager Implement
	manager:= middleware.NewManager() 
	manager.Use(middleware.Logger, middleware.Hudai)

	mux.Handle("GET /", manager.With(
		http.HandlerFunc(handlers.GetRoot),middleware.ArektaHudai))


	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductByID)))

	port := ":8080"
	globalRouter := global_router.GlobalRouter(mux)
	println("🚀 Server is running at http://localhost" + port)

	err := http.ListenAndServe(port, globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}