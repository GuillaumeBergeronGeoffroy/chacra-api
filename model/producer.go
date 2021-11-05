package model

import "time"

// Producer model exportable
type Producer struct {
	ProducerId        uint32
	ProducerEmail     string
	ProducerPassword  string
	ProducerCreatedAt time.Time
	ProducerStatus    uint8
}

// ProducerValidate exportable
func (m Producer) Validate() (err error) {
	return
}
