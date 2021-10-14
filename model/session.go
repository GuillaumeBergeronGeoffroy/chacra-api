package model

// Session model exportable
type Session struct {
	SessionId  string
	UserId     uint32
	ProducerId uint32
	EmployeeId uint32
}

// SessionValidate exportable
func (m Session) Validate() (err error) {
	return
}
