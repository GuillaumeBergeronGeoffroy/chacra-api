package service

import (
	"database/sql"
	"sync"
)

type transactionManager struct {
	db      *sql.DB
	gateway map[string]string
}

var tmOnce sync.Once
var tm transactionManager

// TransactionManager exportable singleton
func TransactionManager(db *sql.DB, gateway map[string]string) transactionManager {
	tmOnce.Do(func() {
		tm = transactionManager{db, gateway}
	})
	return tm
}

// TransactionManagerActions exportable
func (m transactionManager) Actions() (ac Actions, err error) {
	return
}
