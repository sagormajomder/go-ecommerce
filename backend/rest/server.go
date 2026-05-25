package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config){
	// mux.Handle("GET /", middlewares.Hudai(middlewares.Logger(http.HandlerFunc(handlers.GetRoot)))) 
	//* First Manager Implement 
	// manager:= middlewares.NewManager()
	// mux.Handle("GET /", manager.With(
	// 										middlewares.Hudai,
	// 										middlewares.Logger,
	// 										)(http.HandlerFunc(handlers.GetRoot))) 

	//* Second Manager Implement
	// manager:= middlewares.NewManager()
	// mux.Handle("GET /", manager.With(
	// 										http.HandlerFunc(handlers.GetRoot),
	// 										middlewares.Logger, middlewares.Hudai))

	//* Best Manager Implement
	manager:= middlewares.NewManager() 
	manager.Use(middlewares.Logger,middlewares.Cors,middlewares.Preflight)


	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)
	initRoutes(mux, manager)

	addr:= ":"+ strconv.Itoa(cnf.HttpPort)
	println("🚀 Server is running at http://localhost" +addr)

	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}