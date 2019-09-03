package router

import (
	"net/http"

	. "../controller"
	)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Auth     	bool
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index, false},

	// userAccess
	Route{"Login", "POST", "/login", Login, false},

	// webSocket
	Route{"Ws", "GET", "/ws", HandleConnections, true},

	// user
	Route{"IndexUser", "GET", "/user", IndexUser, true},
	Route{"ShowUser", "GET", "/user/{userId}", ShowUser, true},
	Route{"CreateUser", "POST", "/user", CreateUser, true},
	Route{"UpdateUser", "PUT", "/user/{userId}", UpdateUser, true},
	Route{"DeleteUser", "DELETE", "/user/{userId}", DeleteUser, true},

	// product
	Route{"IndexProduct", "GET", "/product", IndexProduct, false},
	Route{"ShowProduct", "GET", "/product/{productId}", ShowProduct, true},
}
