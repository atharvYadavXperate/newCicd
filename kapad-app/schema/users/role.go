package users

type role string

var (
	Admin role = "admin"
	User  role = "user"
)

func (r role) IsValid() bool {
	switch r {
	case Admin, User:
		return true
	default:
		return false
	}
}
