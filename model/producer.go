package model

// Producer model exportable
type Producer struct {
	ProducerId        uint32
	ProducerEmail     string
	ProducerPassword  string
	ProducerName      string
	ProducerCreatedAt string
	ProducerStatus    uint8
}

// ProducerAfterSave exportable
func (m Producer) AfterSave() (err error) {
	return
}

// ProducerBeforeSave exportable
func (m Producer) BeforeSave() (err error) {
	return
}

// ProducerValidate exportable
func (m Producer) Validate() (err error) {
	return
}
