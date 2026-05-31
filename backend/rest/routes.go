package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	// Product Route
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))

	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct), middlewares.AuthenticateJWT))

	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProduct)))

	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(handlers.UpdateProduct), middlewares.AuthenticateJWT))

	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(handlers.DeleteProduct), middlewares.AuthenticateJWT))

	// User Route
	mux.Handle("GET /users", manager.With(http.HandlerFunc(handlers.GetUsers)))
	mux.Handle("POST /users", manager.With(http.HandlerFunc(handlers.CreateUser)))

	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(handlers.Login)))

	mux.Handle("GET /", manager.With(
		http.HandlerFunc(handlers.GetRoot)))

}
