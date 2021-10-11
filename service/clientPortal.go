package service

import (
	"database/sql"
)

// ClientPortal exportable
type ClientPortal struct {
	DB *sql.DB
}

// ClientPortalActionS exportable
func (m ClientPortal) Actions() (ac Actions, err error) {
	return
}
