package model

// OrderFullfilment model exportable
type OrderFullfilment struct {
	OrderFullfilmentId           uint32
	OrderFullfilmentOptionTypeId uint8
	OrderId                      uint32
	OrderFullfilmentStatus       uint8
}

// OrderFullfilmentValidate exportable
func (m OrderFullfilment) Validate() (err error) {
	return
}
