package users

import (
	"io"
	"strings"
	"time"

	"github.com/atharvYadavXperate/newCicd/kapad-app/domain"
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
	u.UserName = strings.TrimSpace(u.UserName)
	u.Password = strings.TrimSpace(u.Password)
}

func (u *UserSchema) IsAllRequiredFields() bool {
	u.TrimSpacesOfUsernamePassword()
	if u.UserName == "" || u.Password == "" || !u.Role.IsValid() {
		return false
	}
	return true
}
