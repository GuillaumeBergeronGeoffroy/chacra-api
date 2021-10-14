package model

// ContentSuggestion model exportable
type ContentSuggestion struct {
	ContentSuggestionId    uint32
	ContentId              uint32
	ContentSuggestionValue string
}

// ContentSuggestionValidate exportable
func (m ContentSuggestion) Validate() (err error) {
	return
}
