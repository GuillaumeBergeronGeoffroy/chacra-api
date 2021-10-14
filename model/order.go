package model

// Order model exportable
type Order struct {
	OrderId     uint32
	OroductId   uint32
	OrderStatus uint8
}

// OrderValidate exportable
func (m Order) Validate() (err error) {
	return
}
