package authfunction

import (
	"net/http"
	"time"

	"github.com/atharvYadavXperate/newCicd/kapad-app/database/users"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user users.UserSchema

	if err := user.ParseData(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Role = "user"
	user.Verified = false
	user.CreatedAt = time.Now()
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := user.ToJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
