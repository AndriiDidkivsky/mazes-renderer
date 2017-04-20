package cors

import (
	"net/http"
)

type Cors struct {
}

func (c *Cors) handlePreflightRequest(w http.ResponseWriter, r *http.Request) {
	//handle origin
	//handle credentials
	//handle methods
	//handle allowed headers
	//handle max age
	//handle exposed headerd
}

func (c *Cors) handleActualRequest(w http.ResponseWriter, r *http.Request) {
	//handle origin
	//handle credentials
	//handle exposed headerd
}

//middleware function constructor
func (c *Cors) Handler(h http.Handler) http.Handler {
	//check request method and delegate to request handler
}
