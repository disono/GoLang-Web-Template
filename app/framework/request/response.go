package request

import (
	"net/http"
	"encoding/json"
)

func RedirectBack(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
}

func JSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func IsMethod(r *http.Request, m string) (bool) {
	return r.Method == m
}