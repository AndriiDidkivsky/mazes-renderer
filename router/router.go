package router

import (
	"net/http"
)

type Route struct {
	handler http.Handler
	method  string
	path    string
}

type Router struct {
	handlers map[string][]http.Handler
	NotFound http.Handler
}

func New() *Router {
	return &Router{}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func (router *Router) GET(path string, handler http.Handler) {
	router.Handle("GET", path, handler)
}

func (router *Router) POST(path string, handler http.Handler) {
	router.Handle("POST", path, handler)
}

func (router *Router) PUT(path string, handler http.Handler) {
	router.Handle("PUT", path, handler)
}

func (router *Router) PATCH(path string, handler http.Handler) {
	router.Handle("PATCH", path, handler)
}

func (router *Router) DELETE(path string, handler http.Handler) {
	router.Handle("DELETE", path, handler)
}

func (router *Router) Handle(method string, path string, handler http.Handler) {

}
