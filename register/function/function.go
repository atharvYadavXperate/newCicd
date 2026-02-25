package authfunction

import (
	"net/http"

	customerror "github.com/atharvYadavXperate/kapad-app/domain/errors"
	"github.com/atharvYadavXperate/kapad-app/schema/users"
	"github.com/atharvYadavXperate/newCicd/kapad-app/domain/helpers"
	"github.com/atharvYadavXperate/newCicd/register/applayer"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, customerror.ErrMethodNotAllows.Error(), http.StatusBadRequest)
		return
	}
	app := applayer.Init()

	var user users.UserSchema

	if err := user.ParseData(r.Body); err != nil {
		http.Error(w, customerror.ErrParseError.Error(), http.StatusBadRequest)
		return
	}

	if !user.IsAllRequiredFields() {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	_, err := app.Database.CreateUser(user)
	if err != nil {
		helpers.ResponseWriterWithError(w, http.StatusBadRequest, "", err)
		return
	}
	helpers.ResponseJsonWriter(w, user, http.StatusCreated, "User created successfully", "")
}
