package service

import "database/sql"

// TransactionManager exportable
type TransactionManager struct {
	DB *sql.DB
}

// ProducerPortalActions exportable
func (m TransactionManager) Actions() (ac Actions, err error) {
	return
}
