package service

import (
	"sync"
)

type transactionManager struct {
	Dao *Dao
}

var tmOnce sync.Once
var tm transactionManager

// TransactionManager exportable singleton
func TransactionManager(dao *Dao) *transactionManager {
	tmOnce.Do(func() {
		tm = transactionManager{dao}
	})
	return &tm
}

// TransactionManagerActions exportable
func (m transactionManager) Actions() (ac Actions, err error) {
	return
}
