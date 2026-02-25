package authfunction

import (
	"log"
	"net/http"
	"time"

	customerror "github.com/atharvYadavXperate/kapad-app/domain/errors"
	"github.com/atharvYadavXperate/kapad-app/domain/helpers"
	"github.com/atharvYadavXperate/kapad-app/schema/users"
	"github.com/atharvYadavXperate/newCicd/register/applayer"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, customerror.ErrMethodNotAllows.Error(), http.StatusBadRequest)
		return
	}
	log.Printf(time.Now().String())
	app := applayer.Init()
	log.Printf(time.Now().String())
	var user users.UserSchema

	if err := user.ParseData(r.Body); err != nil {
		helpers.ResponseWriterWithError(w, http.StatusBadRequest, customerror.ErrParseError.Error(), "")
		return
	}

	if !user.IsAllRequiredFields() {
		helpers.ResponseWriterWithError(w, http.StatusBadRequest, "All fields are required", "")
		return
	}

	if err := user.Validate(); err != nil {
		log.Printf(err.Error())
		helpers.ResponseWriterWithError(w, http.StatusBadRequest, "failed to create account", err.Error())
		return
	}

	_, err := app.Database.CreateUser(user)
	if err != nil {
		log.Printf(err.Error())
		helpers.ResponseWriterWithError(w, http.StatusBadRequest, "username already exits", err.Error())
		return
	}
	log.Printf(time.Now().String())
	helpers.ResponseJsonWriter(w, user, http.StatusCreated, "User created successfully", "")
}
