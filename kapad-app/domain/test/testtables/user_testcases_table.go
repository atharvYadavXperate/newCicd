package testtables

import (
	"github.com/atharvYadavXperate/newCicd/kapad-app/schema/users"
)

type AllFieldsTestCase struct {
	Name        string
	Input       users.UserSchema
	Expected    bool
	ExpectError bool
}

func IsAllRequiredFieldsTable() []AllFieldsTestCase {
	return []AllFieldsTestCase{
		{
			Name: "Valid admin user",
			Input: users.UserSchema{
				UserName: "alice",
				Password: "password123",
				Role:     users.Admin,
			},
			Expected: true,
		},
		{
			Name: "Valid normal user",
			Input: users.UserSchema{
				UserName: "bob",
				Password: "pass123",
				Role:     users.User,
			},
			Expected: true,
		},
		{
			Name: "Valid role with uppercase",
			Input: users.UserSchema{
				UserName: "charlie",
				Password: "mypassword",
				Role:     users.User,
			},
			Expected: true, // depends if your function is case-insensitive
		},
		// Invalid userNames
		{
			Name: "Empty userName",
			Input: users.UserSchema{
				UserName: "   ",
				Password: "pass123",
				Role:     users.User,
			},
			Expected: false,
		},
		{
			Name: "Empty password",
			Input: users.UserSchema{
				UserName: "alice",
				Password: "   ",
				Role:     users.Admin,
			},
			Expected: false,
		},
		{
			Name: "Empty role",
			Input: users.UserSchema{
				UserName: "alice",
				Password: "password123",
				Role:     "",
			},
			Expected: false,
		},
		// Invalid roles
		{
			Name: "Invalid role",
			Input: users.UserSchema{
				UserName: "alice",
				Password: "password123",
				Role:     "superuser",
			},
			Expected: false,
		},
		{
			Name: "Role with spaces",
			Input: users.UserSchema{
				UserName: "alice",
				Password: "password123",
				Role:     " admin ",
			},
			Expected: false,
		},
		// Multiple invalid fields
		{
			Name: "All fields invalid",
			Input: users.UserSchema{
				UserName: "   ",
				Password: "",
				Role:     "manager",
			},
			Expected: false,
		},
		{
			Name: "UserName empty, role invalid",
			Input: users.UserSchema{
				UserName: "",
				Password: "password123",
				Role:     "guest",
			},
			Expected: false,
		},
		{
			Name: "Password empty, role invalid",
			Input: users.UserSchema{
				UserName: "alice",
				Password: "",
				Role:     "guest",
			},
			Expected: false,
		},
		// Very long userName and password
		{
			Name: "Long userName and password",
			Input: users.UserSchema{
				UserName: "user_" + string(make([]byte, 100)),
				Password: string(make([]byte, 200)),
				Role:     users.User,
			},
			Expected: true,
		},
	}
}

type TrimSpaceCases struct {
	Name             string
	Input            users.UserSchema
	ExpectedUserName string
	ExpectedPassword string
}

func TrimSpacesOfUserNamePasswordTable() []TrimSpaceCases {
	return []TrimSpaceCases{
		{
			Name: "Valid admin user",
			Input: users.UserSchema{
				UserName: "alice",
				Password: "password123",
				Role:     users.Admin,
			},
			ExpectedUserName: "alice",
			ExpectedPassword: "password123",
		},
		{
			Name: "Valid normal user",
			Input: users.UserSchema{
				UserName: "bob",
				Password: "pass123",
				Role:     users.User,
			},
			ExpectedUserName: "bob",
			ExpectedPassword: "pass123",
		},
		{
			Name: "Valid role with uppercase",
			Input: users.UserSchema{
				UserName: "charlie",
				Password: "mypassword",
				Role:     users.User,
			},
			ExpectedUserName: "charlie",
			ExpectedPassword: "mypassword", // depends if your function is case-insensitive
		},
		// Invalid userNames
		{
			Name: "Empty userName",
			Input: users.UserSchema{
				UserName: "   ",
				Password: "pass123",
				Role:     users.User,
			},
			ExpectedUserName: "",
			ExpectedPassword: "pass123",
		},
		{
			Name: "Empty password",
			Input: users.UserSchema{
				UserName: "alice",
				Password: "   ",
				Role:     users.Admin,
			},
			ExpectedUserName: "alice",
			ExpectedPassword: "",
		},
		{
			Name: "Empty role",
			Input: users.UserSchema{
				UserName: "alice",
				Password: "password123",
				Role:     "",
			},
			ExpectedUserName: "alice",
			ExpectedPassword: "password123",
		},
		// Invalid roles
		{
			Name: "Invalid role",
			Input: users.UserSchema{
				UserName: "alice",
				Password: "password123",
				Role:     "superuser",
			},
			ExpectedUserName: "alice",
			ExpectedPassword: "password123",
		},
		{
			Name: "Role with spaces",
			Input: users.UserSchema{
				UserName: "alice",
				Password: "password123",
				Role:     " admin ",
			},
			ExpectedUserName: "alice",
			ExpectedPassword: "password123",
		},
		// Multiple invalid fields
		{
			Name: "All fields invalid",
			Input: users.UserSchema{
				UserName: "   ",
				Password: "",
				Role:     "manager",
			},
			ExpectedUserName: "",
			ExpectedPassword: "",
		},
	}
}

type ToJSONStringCase struct {
	Name        string
	Input       users.UserSchema
	ExpectError bool
}

func ToJSONStringTable() []ToJSONStringCase {
	return []ToJSONStringCase{
		{
			Name: "Normal user",
			Input: users.UserSchema{
				UserName: "alice",
				Password: "password123",
				Role:     "admin",
			},
			ExpectError: false,
		},
		{
			Name: "Empty fields",
			Input: users.UserSchema{
				UserName: "",
				Password: "",
				Role:     "",
			},
			ExpectError: false,
		},
	}
}

type UsernameTestCase struct {
	Input    string
	Expected bool
}

func ValidateUsernameTable() []UsernameTestCase {
	return []UsernameTestCase{
		{"john", true},
		{"  alice  ", true}, // spaces should be trimmed
		{"ab", false},       // too short
		{"", false},         // empty
	}
}

type PasswordTestCase struct {
	Input    string
	Expected bool
}

func ValidatePasswordTable() []PasswordTestCase {
	return []PasswordTestCase{
		{"password1", true},
		{"pass1", false},
		{"password", false},
		{"12345678", false},    // no letter
		{"  p4ssword  ", true}, // spaces should be trimmed
		{"", false},            // empty
	}
}
