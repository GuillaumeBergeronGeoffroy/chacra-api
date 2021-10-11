package models

import "time"

// ProductAvailability model exportable
type ProductAvailability struct {
	productAvailabilityId       uint32
	productId                   uint32
	productAvailabilityQuantity uint16
	productAvailabilityStart    time.Time
	productAvailabilityEnd      time.Time
}
