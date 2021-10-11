package models

// OrderFullfilment model exportable
type OrderFullfilment struct {
	orderFullfilmentId           uint32
	orderFullfilmentOptionTypeId uint8
	orderId                      uint32
	orderFullfilmentStatus       uint8
}
