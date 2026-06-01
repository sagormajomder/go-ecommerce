package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle("GET /", manager.With(
		http.HandlerFunc(handlers.GetRoot)))

}
