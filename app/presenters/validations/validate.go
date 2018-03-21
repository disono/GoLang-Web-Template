package validations

import (
	"net/http"
	"github.com/thedevsaddam/govalidator"
	"github.com/disono/GoLang-Web-Template/app/framework/request"
)

func options(r *http.Request, rules map[string][]string) (govalidator.Options) {
	return govalidator.Options{
		Request:         r,
		Rules:           rules,
		RequiredDefault: true,
	}
}

func Handle(rules map[string][]string, w http.ResponseWriter, r *http.Request) (bool) {
	v := govalidator.New(options(r, rules))
	e := v.Validate()

	if len(e) > 0 {
		err := map[string]interface{}{"success": false, "errors": e}
		request.JSONResponse(w, err)

		return false
	} else {
		return true
	}
}
