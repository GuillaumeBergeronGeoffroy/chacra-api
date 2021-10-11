package service

import "database/sql"

// ProducerPortal exportable
type ProducerPortal struct {
	DB *sql.DB
}

// ProducerPortalActions exportable
func (m ProducerPortal) Actions() (ac Actions, err error) {
	return
}
