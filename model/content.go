package models

// Content model exportable
type Content struct {
	ContentId          uint32
	ContentModel       string
	ContentModelId     uint32
	ContentModelTypeId uint8
	ContentLang        string
	ContentValue       string
	ContentStatus      uint8
}
