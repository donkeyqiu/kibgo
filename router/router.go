package router

import (
	"net/http"

	. "../middleware"

	"github.com/gorilla/mux"
	)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	router := r.StrictSlash(true).PathPrefix("/api/v1/").Subrouter()
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		// on/off Auth Middleware
		handler = Auth(handler, route.Name, route.Auth)
		// handler = ServeHTTP(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
