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
		InitServiceSqlDB(ep.Dao.DB, epInitSql)
	})
	return &ep
}

// EmployeePortalActions exportable
func (m employeePortal) Actions() (ac Actions, err error) {
	return
}

var epInitSql = []string{
	`CREATE TABLE Employee (
		EmployeeId INT NOT NULL AUTO_INCREMENT,
		EmployeeEmail VARCHAR(255) NOT NULL,
		EmployeePassword VARCHAR(255) NOT NULL,
		PRIMARY KEY (EmployeeId),
		CONSTRAINT uidx_Employee_EmployeeEmail UNIQUE (EmployeeEmail)
	);`,
}
