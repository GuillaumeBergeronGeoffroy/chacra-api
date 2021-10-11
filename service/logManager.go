package service

import "database/sql"

// LogManager exportable
type LogManager struct {
	DB    *sql.DB
	links map[string]string
}

// LogManagerActions exportable
func (m LogManager) Actions() (ac Actions, err error) {
	return
}
