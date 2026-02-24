package kapadtest

import (
	"testing"

	"github.com/atharvYadavXperate/newCicd/kapad-app/domain/test/testtables"
	"github.com/atharvYadavXperate/newCicd/kapad-app/schema/users"
)

func TestUserSchema_AllFieldsRequired(t *testing.T) {
	cases := testtables.IsAllRequiredFieldsTable()
	for _, c := range cases {
		result := c.Input.IsAllRequiredFields()
		if result != c.Expected {
			t.Errorf("expected %v, got %v", c.Expected, result)
		}
	}
}

func TestUserSchema_TrimSpacesOfUsernamePassword(t *testing.T) {
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

func TestUserSchema_ToJSONString(t *testing.T) {
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

func TestUserSchema_ValidateUsernamePassword(t *testing.T) {
	// Username test cases
	for _, c := range testtables.ValidateUsernameTable() {
		u := users.UserSchema{UserName: c.Input, Password: "validPass1", Role: users.User}
		err := u.Validate()
		got := err == nil
		if got != c.Expected {
			t.Errorf("Test %q: Validate username %q expected valid=%v, got %v", c.Input, c.Input, c.Expected, got)
		}
	}

	// Password test cases
	for _, c := range testtables.ValidatePasswordTable() {
		u := users.UserSchema{UserName: "validUser", Password: c.Input, Role: users.User}
		err := u.Validate()
		got := err == nil
		if got != c.Expected {
			t.Errorf("Test %q: Validate password %q expected valid=%v, got %v", c.Input, c.Input, c.Expected, got)
		}
	}
}
