package middlewares

import (
	"net/http"
	"github.com/disono/GoLang-Web-Template/app/framework/request"
)

// Logging logs all requests with its path and the time it took to process
func VerifyCSRFToken() request.Middleware {
	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			if request.IsMethod(r, "POST") {

			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}
