package model

import "time"

// ProductAvailability model exportable
type ProductAvailability struct {
	ProductAvailabilityId       uint32
	ProductId                   uint32
	ProductAvailabilityQuantity uint16
	ProductAvailabilityStart    time.Time
	ProductAvailabilityEnd      time.Time
}

// ProductAvailabilityValidate exportable
func (m ProductAvailability) Validate() (err error) {
	return
}
