package users

import (
	"time"
)

type UserSchema struct {
	UserName  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"createdAt"`
}
