package kapadtest

import (
	"testing"

	"github.com/atharvYadavXperate/newCicd/kapad-app/domain/test/testtables"
)

func TestAllFieldsRequired(t *testing.T) {
	cases := testtables.IsAllRequiredFieldsTable()
	for _, c := range cases {
		result := c.Input.IsAllRequiredFields()
		if result != c.Expected {
			t.Errorf("expected %v, got %v", c.Expected, result)
		}
	}
}

func TestTrimSpacesOfUsernamePassword(t *testing.T) {
	cases := testtables.TrimSpacesOfUserNamePasswordTable()
	for _, c := range cases {
		c.Input.TrimSpacesOfUsernamePassword()
		if c.Input.UserName != c.ExpectedUserName {
			t.Errorf("expected username %v, got %v", c.ExpectedUserName, c.Input.UserName)
		}
		if c.Input.Password != c.ExpectedPassword {
			t.Errorf("expected password %v, got %v", c.ExpectedUserName, c.Input.UserName)
		}
	}
}

func TestToJSONString(t *testing.T) {
	cases := testtables.ToJSONStringTable()
	for _, c := range cases {
		_, err := c.Input.ToJSONString()
		if c.ExpectError {
			if err == nil {
				t.Errorf("expected an error but got nil")
			}
			return
		}
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
	}
}
