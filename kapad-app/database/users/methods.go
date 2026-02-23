package users

import (
	"io"

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
	str, err := domain.ToJSONString(*u)
	return str, err
}
