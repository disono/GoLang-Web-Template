package authentication

import (
	"net/http"
	"fmt"
)

func RegisterViewAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Register View: %s\n", r.URL.Path)
}

func RegisterStoreAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Register Store: %s\n", r.URL.Path)
}