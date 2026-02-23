package users

import (
	"time"
)

type UserSchema struct {
	username  string
	password  string
	role      role
	verified  bool
	createdAt time.Time
}
