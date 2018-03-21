package web

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/disono/GoLang-Web-Template/app/framework/request"
)

type PageData struct {
	Title string
}

func ShowAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := PageData{
		Title: vars["slug"],
	}

	request.View(w, "pages/show", data)
}

func ListAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := PageData{
		Title: vars["name"],
	}

	request.View(w, "pages/list", data)
}
