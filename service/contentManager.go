package service

import (
	"database/sql"
)

// ContentManager exportable
type ContentManager struct {
	DB    *sql.DB
	links map[string]string
}

// ContentManagerActions exportable
func (m ContentManager) Actions() (ac Actions, err error) {
	return
}
