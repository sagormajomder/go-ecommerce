package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve(){
	port := ":8080"
	
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
	manager.Use(middleware.Logger,middleware.Cors,middleware.Preflight)


	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)
	initRoutes(mux, manager)

	println("🚀 Server is running at http://localhost" + port)

	err := http.ListenAndServe(port, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}