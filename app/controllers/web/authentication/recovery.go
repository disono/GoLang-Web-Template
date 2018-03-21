package authentication

import (
	"net/http"
	"fmt"
)

func ForgotViewAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Forgot View: %s\n", r.URL.Path)
}

func ForgotStoreAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Forgot Store: %s\n", r.URL.Path)
}

func VerifyForgotView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Verify Forgot View: %s\n", r.URL.Path)
}

func ChangePasswordView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Verify Forgot View: %s\n", r.URL.Path)
}