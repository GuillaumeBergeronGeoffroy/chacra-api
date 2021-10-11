package models

// Content model exportable
type Content struct {
	contentId          uint32
	contentModel       string
	contentModelId     uint32
	contentModelTypeId uint8
	contentLang        string
	contentValue       string
	contentStatus      uint8
}
