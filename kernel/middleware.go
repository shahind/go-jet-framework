package kernel

import (
	"github.com/gorilla/mux"
	"github.com/shahind/go-jet-framework/register"
)

// Parse list of middleware and get an array of []mux.Middleware func
// Required by Gorilla Mux
func parseMiddleware(mwList []register.Middleware) []mux.MiddlewareFunc {
	var midFunc []mux.MiddlewareFunc

	for i := len(mwList) - 1; i > -1; i-- {
		midFunc = append(midFunc, mwList[i].Handle)
	}

	return midFunc
}
