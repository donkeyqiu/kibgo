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
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},

	// userAccess
	Route{"Login", "POST", "/login", Login},

	// user
	Route{"IndexUser", "GET", "/user", IndexUser},
	Route{"ShowUser", "GET", "/user/{userId}", ShowUser},
	Route{"CreateUser", "POST", "/user", CreateUser},
	Route{"UpdateUser", "PUT", "/user/{userId}", UpdateUser},
	Route{"DeleteUser", "DELETE", "/user/{userId}", DeleteUser},

	// product
	Route{"IndexProduct", "GET", "/product", IndexProduct},
	Route{"ShowProduct", "GET", "/product/{productId}", ShowProduct},
}
