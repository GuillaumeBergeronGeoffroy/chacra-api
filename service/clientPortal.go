package service

import (
	"sync"
)

type clientPortal struct {
	Dao *Dao
}

var cpOnce sync.Once
var cp clientPortal

// ClientPortal singleton exportable
func ClientPortal(dao *Dao) *clientPortal {
	cpOnce.Do(func() {
		cp = clientPortal{dao}
		InitServiceSqlDB(cp.Dao.DB, cpInitSql)
	})
	return &cp
}

// ClientPortalActions exportable
func (m clientPortal) Actions() (ac Actions, err error) {
	return
}

var cpInitSql = []string{
	`CREATE TABLE User (
		UserId INT NOT NULL AUTO_INCREMENT,
		UserEmail VARCHAR(255) NOT NULL,
		UserPassword VARCHAR(255) NOT NULL,
		PRIMARY KEY (UserId),
		CONSTRAINT uidx_User_UserEmail UNIQUE (UserEmail)
	);`,
}
