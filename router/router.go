package router

import (
	"fmt"
	"net/http"
)

type Route struct {
	handler http.HandlerFunc
	method  string
	path    string
}

type Router struct {
	handlers map[string][]*Route
	NotFound http.Handler
}

func New() *Router {
	return &Router{handlers: make(map[string][]*Route)}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handlers, ok := router.handlers[r.Method]; ok {
		for _, handler := range handlers {
			fmt.Println(handler.path, r.URL.Path[1:])
			//todo: develop path parse logic and routing tree
			if handler.path == r.URL.Path[1:] {
				handler.handler(w, r)
			}
		}
	}
}

func (router *Router) GET(path string, handler http.HandlerFunc) {
	router.Handle("GET", path, handler)
}

func (router *Router) POST(path string, handler http.HandlerFunc) {
	router.Handle("POST", path, handler)
}

func (router *Router) PUT(path string, handler http.HandlerFunc) {
	router.Handle("PUT", path, handler)
}

func (router *Router) PATCH(path string, handler http.HandlerFunc) {
	router.Handle("PATCH", path, handler)
}

func (router *Router) DELETE(path string, handler http.HandlerFunc) {
	router.Handle("DELETE", path, handler)
}

func (router *Router) Handle(method string, path string, handler http.HandlerFunc) {
	if router.handlers[method] == nil {
		router.handlers[method] = append(router.handlers[method], &Route{method: method, path: path, handler: handler})
	}
}
