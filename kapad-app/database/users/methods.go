package users

import (
	"io"
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
	if u.Role == "" {
		u.Role = "user" // default role
	}
	u.Verified = false
	u.CreatedAt = time.Now()
}
