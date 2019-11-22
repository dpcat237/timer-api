package model

import (
	"net/http"
)

// Route defines properties of the HTTP route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRoute(m, p string, h http.HandlerFunc, n string) Route {
	return Route{
		Name:        n,
		Method:      m,
		Pattern:     p,
		HandlerFunc: h,
	}
}
