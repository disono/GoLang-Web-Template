package authentication

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/disono/GoLang-Web-Template/app/framework/request"
	"github.com/disono/GoLang-Web-Template/app/presenters/validations/web/authentications"
	"fmt"
)

type PageData struct {
	Title string
}

func LoginViewAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := PageData{
		Title: vars["slug"],
	}

	request.View(w, "authentications/login", data)
}

func AuthenticateAction(w http.ResponseWriter, r *http.Request) {
	if authentications.Run(w, r) {
		fmt.Print(r.FormValue("username"))
		request.RedirectBack(w, r)
	}
}
