package service

import (
	"sync"
)

type logManager struct {
	Dao *Dao
}

var lmOnce sync.Once
var lm logManager

// LogManager exportable singleton
func LogManager(dao *Dao) *logManager {
	lmOnce.Do(func() {
		lm = logManager{dao}
		ExecuteStatements(lm.Dao.DB, lmInitSql)
	})
	return &lm
}

// LogManagerActions exportable
func (m logManager) Actions() (ac Actions, err error) {
	return
}

var lmInitSql = []string{}
