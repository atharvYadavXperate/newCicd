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

var app = applayer.Init() // ✅ initialize once (IMPORTANT)

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, customerror.ErrMethodNotAllows.Error(), http.StatusMethodNotAllowed)
		return
	}

	log.Printf("Request received at %v", time.Now())

	var user users.UserSchema
	var usersList []users.UserSchema
	// Send multiple users as response
	helpers.ResponseJsonWriter(w, usersList, http.StatusCreated, "Users created successfully", "")
	if err := user.ParseData(r.Body); err != nil {
		helpers.ResponseWriterWithError(w, http.StatusBadRequest, customerror.ErrParseError.Error(), "")
		return
	}

	if !user.IsAllRequiredFields() {
		helpers.ResponseWriterWithError(w, http.StatusBadRequest, "All fields are required", "")
		return
	}

	if err := user.Validate(); err != nil {
		log.Println("Validation error:", err)
		helpers.ResponseWriterWithError(w, http.StatusBadRequest, "Failed to create account", err.Error())
		return
	}

	log.Printf("Storing in database at %v", time.Now())

	_, err := app.Database.CreateUser(user)
	if err != nil {
		log.Println("Database error:", err)
		helpers.ResponseWriterWithError(w, http.StatusBadRequest, "Username already exists", err.Error())
		return
	}
	for i := 0; i < 50; i++ {
		usersList = append(usersList, user)
	}
	helpers.ResponseJsonWriter(w, usersList, http.StatusCreated, "User created successfully", "")
}
