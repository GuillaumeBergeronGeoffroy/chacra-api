package service

import (
	"sync"
)

type employeePortal struct {
	Dao *Dao
}

var epOnce sync.Once
var ep employeePortal

// EmployeePortal exportable singleton
func EmployeePortal(dao *Dao) *employeePortal {
	epOnce.Do(func() {
		ep = employeePortal{dao}
	})
	return &ep
}

// EmployeePortalActions exportable
func (m employeePortal) Actions() (ac Actions, err error) {
	return
}
