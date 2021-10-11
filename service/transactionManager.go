package service

import "database/sql"

// TransactionManager exportable
type TransactionManager struct {
	DB    *sql.DB
	links map[string]string
}

// ProducerPortalActions exportable
func (m TransactionManager) Actions() (ac Actions, err error) {
	return
}
