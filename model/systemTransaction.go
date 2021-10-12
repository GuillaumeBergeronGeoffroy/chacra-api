package models

import "time"

// SystemTransaction model exportable
type SystemTransaction struct {
	SystemTransactionId        uint32
	SystemTransactionTypeId    uint8
	UserId                     uint32
	ProducerId                 uint32
	EmployeeId                 uint32
	SystemTransactionCreatedAt time.Time
}
