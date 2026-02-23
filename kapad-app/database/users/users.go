package users

import (
	"time"
)

type UserSchema struct {
	UserName  string
	Password  string
	Role      role
	Verified  bool
	CreatedAt time.Time
}
