package users

import (
	"io"
	"time"

	"github.com/atharvYadavXperate/newCicd/kapad-app/domain"
	"github.com/atharvYadavXperate/newCicd/kapad-app/domain/helpers"
)

func (u *UserSchema) ParseData(body io.Reader) error {
	data, err := domain.DecodeBody[UserSchema](body)
	if err != nil {
		return err
	}
	*u = *data
	return nil
}

func (u *UserSchema) ToJSON() ([]byte, error) {
	return domain.ToJSON(*u)
}

func (u *UserSchema) ToJSONString() (string, error) {
	return domain.ToJSONString(*u)
}

func (u *UserSchema) SetDefaults() {
	u.Verified = false
	u.CreatedAt = time.Now()
}

func (u *UserSchema) TrimSpacesOfUsernamePassword() {
	u.UserName = helpers.TrimSpacesInString(u.UserName)
	u.Password = helpers.TrimSpacesInString(u.Password)
}

func (u *UserSchema) IsAllRequiredFields() bool {
	if helpers.IsEmptyString(u.UserName) || helpers.IsEmptyString(u.Password) || !u.Role.IsValid() {
		return false
	}
	return true
}

func (u *UserSchema) Validate() error {
	u.TrimSpacesOfUsernamePassword()
	if err := helpers.ValidateUsername(u.UserName); err != nil {
		return err
	}
	if err := helpers.ValidatePassword(u.Password); err != nil {
		return err
	}
	return nil
}

func (u *UserSchema) HashPassword() {
	u.Password = helpers.HashValueUsingBcrypt(u.Password)
}

func (u *UserSchema) IsCorrectPassword(password string) bool {
	return helpers.IsHashedAndPlanStringEqual(u.Password, password)
}
