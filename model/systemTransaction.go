package model

// SystemTransaction model exportable
type SystemTransaction struct {
	SystemTransactionId        uint32
	SystemTransactionTypeId    uint8
	UserId                     uint32
	ProducerId                 uint32
	EmployeeId                 uint32
	SystemTransactionCreatedAt string
}

// SystemTransactionValidate exportable
func (m Session) SystemTransaction() (err error) {
	return
}
