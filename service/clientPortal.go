package service

import (
	"database/sql"
)

// ClientPortal exportable
type ClientPortal struct {
	DB *sql.DB
	links map[string]string
}

// ClientPortalActionS exportable
func (m ClientPortal) Actions() (ac Actions, err error) {
	return
}
