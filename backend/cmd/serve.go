package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Serve(){
	cfn:= config.GetConfig()
	
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

	addr:= ":"+ strconv.Itoa(cfn.HttpPort)
	println("🚀 Server is running at http://localhost" +addr)

	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}