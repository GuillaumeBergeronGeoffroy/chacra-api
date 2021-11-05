package model

import "time"

const status_waitlist = 0
const status_active = 1

// User model exportable
type User struct {
	UserId        uint32
	UserEmail     string
	UserPassword  string
	UserCreatedAt time.Time
	UserStatus    uint8
}

// UserValidate exportable
func (m User) Validate() (err error) {
	return
}
