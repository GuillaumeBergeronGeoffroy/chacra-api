package service

import "database/sql"

// EmployeePortal exportable
type EmployeePortal struct {
	DB *sql.DB
}

// EmployeePortalActions exportable
func (m EmployeePortal) Actions() (ac Actions, err error) {
	return
}