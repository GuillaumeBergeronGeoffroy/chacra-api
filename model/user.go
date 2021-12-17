package model

const status_waitlist = 0
const status_active = 1

// User model exportable
type User struct {
	UserId        uint32
	UserEmail     string
	UserPassword  string
	UserName      string
	UserCreatedAt string
	UserStatus    uint8
}

// UserAfterSave exportable
func (m User) AfterSave() (err error) {
	return
}

// UserBeforeSave exportable
func (m User) BeforeSave() (err error) {
	return
}

// UserValidate exportable
func (m User) Validate() (err error) {
	return
}
