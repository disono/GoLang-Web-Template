package request

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain applies middle wares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middleWares ...Middleware) http.HandlerFunc {
	for _, m := range middleWares {
		f = m(f)
	}

	return f
}
