package models

import "time"

// Producer model
type Producer struct {
	producerId        int32
	producerEmail     string
	producerPassword  string
	producerCreatedAt time.Time
	producerStatus    int8
}
