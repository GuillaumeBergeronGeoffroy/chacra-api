package model

// ProductFullfilmentOption model exportable
type ProductFullfilmentOption struct {
	ProductFullfilmentOptionId     uint32
	ProductId                      uint32
	ProductFullfilmentOptionTypeId uint8
}

// ProductFullfilmentOptionValidate exportable
func (m ProductFullfilmentOption) Validate() (err error) {
	return
}
