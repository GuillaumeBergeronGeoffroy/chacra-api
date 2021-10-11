package models

import "time"

// SystemTransaction model exportable
type SystemTransaction struct {
	systemTransactionId        uint32
	systemTransactionTypeId    uint8
	userId                     uint32
	producerId                 uint32
	employeeId                 uint32
	systemTransactionCreatedAt time.Time
}
