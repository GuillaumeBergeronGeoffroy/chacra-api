package service

import (
	"database/sql"
	"sync"
)

type logManager struct {
	db      *sql.DB
	gateway map[string]string
}

var lmOnce sync.Once
var lm logManager

// LogManager exportable singleton
func LogManager(db *sql.DB, gateway map[string]string) logManager {
	lmOnce.Do(func() {
		lm = logManager{db, gateway}
	})
	return lm
}

// LogManagerActions exportable
func (m logManager) Actions() (ac Actions, err error) {
	return
}
