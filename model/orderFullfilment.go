package models

// OrderFullfilment model exportable
type OrderFullfilment struct {
	OrderFullfilmentId           uint32
	OrderFullfilmentOptionTypeId uint8
	OrderId                      uint32
	OrderFullfilmentStatus       uint8
}
