package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

/*
func (mngr *Manager)With (middlewares ...Middleware)Middleware{

	return func(next http.Handler) http.Handler {
		n:=next
 		// middlewares = [Hudai, Logger]
		for i:=len(middlewares) - 1; i>=0; i--{
			middleware := middlewares[i]
			n = middleware(n)
		}
		return n // middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.GetProducts)))
	}
}
*/

/*
func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler{

	n := next

	// middlewares = [Logger,Hudai]
	for _,middleware := range middlewares{
		n = middleware(n)
	}

	return n  // middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.GetProducts)))
}
*/

// builder pattern
func (mgnr *Manager) Use(middlewares ...Middleware) {
	// middlewares = [Logger,Cors,Preflight]
	mgnr.globalMiddlewares = append(mgnr.globalMiddlewares, middlewares...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {

	n := next

	for _, middleware := range middlewares {
		n = middleware(n)
	}

	return n
}
func (mngr *Manager) WrapMux(next http.Handler) http.Handler {

	n := next

	// mngr.globalMiddlewares = [Preflight,Cors, Logger]
	// for _, globalMiddleware := range mngr.globalMiddlewares{
	// 	n = globalMiddleware(n)
	// }

	// mngr.globalMiddlewares = [Logger,Cors,Preflight]
	for i := len(mngr.globalMiddlewares) - 1; i >= 0; i-- {
		n = mngr.globalMiddlewares[i](n)
	}

	return n // middleware.Logger(middleware.Cors(middleware.Preflight(http.HandlerFunc(mux))))
}
