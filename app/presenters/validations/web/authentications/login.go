package authentications

import (
	"net/http"
	"github.com/thedevsaddam/govalidator"
	"github.com/disono/GoLang-Web-Template/app/presenters/validations"
)

func Run(w http.ResponseWriter, r *http.Request) (bool) {
	return handle(w, r)
}

func rules() (map[string][]string) {
	return govalidator.MapData{
		"username": []string{"required"},
		"password": []string{"required"},
	}
}

func handle(w http.ResponseWriter, r *http.Request) (bool) {
	return validations.Handle(rules(), w, r)
}
