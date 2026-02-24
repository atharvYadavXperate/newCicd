package authfunction

import (
	"net/http"
	"time"

	"github.com/atharvYadavXperate/newCicd/kapad-app/database/users"
	customerror "github.com/atharvYadavXperate/newCicd/kapad-app/domain/errors"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, customerror.ErrMethodNotAllows.Error(), http.StatusBadRequest)
		return
	}
	var user users.UserSchema

	if err := user.ParseData(r.Body); err != nil {
		http.Error(w, customerror.ErrParseError.Error(), http.StatusBadRequest)
		return
	}

	if user.IsAllRequiredFields() {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

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
