package model

// User model exportable
type User struct {
	UserId       uint32
	UserEmail    string
	UserPassword string
}

// UserValidate exportable
func (m User) Validate() (err error) {
	return
}
