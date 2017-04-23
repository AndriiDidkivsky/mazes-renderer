package cors

import (
	"net/http"
	"strings"
)

const (
	accessControlOrigin  = "Access-Control-Allow-Origin"
	accessControlMethods = "Access-Control-Allow-Methods"
	accessControlHeaders = "Access-Control-Allow-Headers"
)

type Options struct {
	AllowOrigins string
	AllowMethods []string
	AllowHeaders []string
}

type Cors struct {
	allowOrigins string
	allowMethods string
	allowHeaders string
}

func concatHeaders(headers []string) string {
	return strings.Join(headers, ", ")
}

func New(o *Options) *Cors {
	return &Cors{
		allowOrigins: o.AllowOrigins,
		allowMethods: concatHeaders(o.AllowMethods),
		allowHeaders: concatHeaders(o.AllowHeaders),
	}
}

func (c *Cors) handlePreflightRequest(w http.ResponseWriter, r *http.Request) {
	//handle origin
	//handle credentials
	//handle methods
	//handle allowed headers
	//handle max age
	//handle exposed headerd
	w.Header().Set(accessControlOrigin, c.allowOrigins)
	w.Header().Set(accessControlMethods, c.allowMethods)
	w.Header().Set(accessControlHeaders, c.allowHeaders)
}

func (c *Cors) handleActualRequest(w http.ResponseWriter, r *http.Request) {
	//handle origin
	//handle credentials
	//handle exposed headerd
	w.Header().Set(accessControlOrigin, c.allowOrigins)

}

//middleware function constructor
func (c *Cors) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			c.handlePreflightRequest(w, r)
		} else {
			c.handleActualRequest(w, r)
		}
		h.ServeHTTP(w, r)
	})
}
