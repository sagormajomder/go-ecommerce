package cmd

import (
	"ecommerce/global_router"
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

	initRoutes(mux, manager)

	port := ":8080"
	globalRouter := global_router.GlobalRouter(mux)
	println("🚀 Server is running at http://localhost" + port)

	err := http.ListenAndServe(port, globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}