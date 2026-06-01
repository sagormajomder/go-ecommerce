package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	userHandler    *user.Handler
	productHandler *product.Handler
}

func NewServer(cnf *config.Config, userHandler *user.Handler, productHandler *product.Handler) *Server {
	return &Server{
		cnf:            cnf,
		userHandler:    userHandler,
		productHandler: productHandler,
	}
}

func (server *Server) Start() {
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
	manager := middlewares.NewManager()
	manager.Use(middlewares.Logger, middlewares.Cors, middlewares.Preflight)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)
	server.productHandler.RegisterHandlers(mux, manager)
	server.userHandler.RegisterHandlers(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	println("🚀 Server is running at http://localhost" + addr)

	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
