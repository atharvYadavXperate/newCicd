package authfunction

import (
	"net/http"

	"github.com/atharvYadavXperate/newCicd/kapad-app/database/users"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user = users.UserSchema{}
	user.ParseData(r.Body)
	w.Header().Set("Content-type", "application/json")
	json, err := user.ToJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(json)
}
