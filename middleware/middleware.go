package middleware

import (
	"net/http"
)

type HandlerConstructor func(http.Handler) http.Handler

type Middleware struct {
	constructors []HandlerConstructor
}

func New(constructors ...HandlerConstructor) *Middleware {
	return &Middleware{constructors}
}

func (m *Middleware) Bind(h http.Handler) http.Handler {
	for i := range m.constructors {
		h = m.constructors[len(m.constructors)-1-i](h)
	}
	return h
}
