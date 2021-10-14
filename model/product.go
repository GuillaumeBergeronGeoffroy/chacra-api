package model

// Product model exportable
type Product struct {
	ProductId     uint32
	ProducerId    uint32
	ProductStatus uint8
}

// ProductValidate exportable
func (m Product) Validate() (err error) {
	return
}
