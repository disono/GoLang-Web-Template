package routes

import (
	"github.com/gorilla/mux"
	"github.com/disono/GoLang-Web-Template/app/controllers/web"
	"github.com/disono/GoLang-Web-Template/app/controllers/web/authentication"
	"net/http"
	"time"
	"log"
	"github.com/disono/GoLang-Web-Template/app/framework/request"
	"github.com/disono/GoLang-Web-Template/app/presenters/middlewares"
)

func InitializeRoutes() {
	r := mux.NewRouter()

	// assets
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("resources/static/public"))))
	r.PathPrefix("/private/").Handler(http.StripPrefix("/private/", http.FileServer(http.Dir("resources/static/private"))))

	// pages
	r.HandleFunc("/", web.ShowAction).Methods("GET")
	r.HandleFunc("/p/{slug}", web.ShowAction).Methods("GET")
	r.HandleFunc("/p/category/{name}", web.ListAction).Methods("GET")
	r.HandleFunc("/p/archive/{year}/{month}", web.ListAction).Methods("GET")

	// authentication
	r.HandleFunc("/login", request.Chain(authentication.LoginViewAction, middlewares.VerifyCSRFToken())).Methods("GET")
	r.HandleFunc("/login", authentication.AuthenticateAction).Methods("POST")
	r.HandleFunc("/register", authentication.RegisterViewAction).Methods("GET")
	r.HandleFunc("/register", authentication.RegisterStoreAction).Methods("POST")
	r.HandleFunc("/forgot", authentication.ForgotViewAction).Methods("GET")
	r.HandleFunc("/forgot", authentication.ForgotStoreAction).Methods("POST")
	r.HandleFunc("/forgot/verify", authentication.VerifyForgotView).Methods("GET")
	r.HandleFunc("/forgot/verify", authentication.ChangePasswordView).Methods("POST")

	serve(r)
}

func serve(r *mux.Router) {
	srv := &http.Server{
		Handler:      r,
		Addr:         ":80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
