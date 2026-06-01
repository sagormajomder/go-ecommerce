package user

import (
	"ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterHandlers(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("GET /users", manager.With(http.HandlerFunc(h.GetUsers)))
	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUser)))

	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(h.Login)))

}
