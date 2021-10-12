package service

import (
	"database/sql"
	"sync"
)

type employeePortal struct {
	db      *sql.DB
	gateway map[string]string
}

var epOnce sync.Once
var ep employeePortal

// EmployeePortal exportable singleton
func EmployeePortal(db *sql.DB, gateway map[string]string) employeePortal {
	epOnce.Do(func() {
		ep = employeePortal{db, gateway}
	})
	return ep
}

// EmployeePortalActions exportable
func (m employeePortal) Actions() (ac Actions, err error) {
	return
}
