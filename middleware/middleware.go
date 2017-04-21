package middleware

import (
	"net/http"
)

type HandlerConstructor func(http.Handler) http.Handler

type Queue struct {
	constructors []HandlerConstructor
}

type Middleware struct {
}

func New(constructors ...HandlerConstructor) Queue {
	return Queue{constructors}
}

func (q Queue) Bind(h http.Handler) http.Handler {
	for i := range q.constructors {
		h = q.constructors[len(q.constructors)-1-i](h)
	}
	return h
}
