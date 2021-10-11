package service

import "database/sql"

// ProducerPortal exportable
type ProducerPortal struct {
	DB    *sql.DB
	links map[string]string
}

// ProducerPortalActions exportable
func (m ProducerPortal) Actions() (ac Actions, err error) {
	return
}
