package model

// Employee model exportable
type Employee struct {
	EmployeeId       uint32
	EmployeeEmail    string
	EmployeePassword string
}

// EmployeeValidate exportable
func (m Employee) Validate() (err error) {
	return
}
