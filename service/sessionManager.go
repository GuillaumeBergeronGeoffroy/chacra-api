package service

import (
	"sync"
)

type sessionManager struct {
	Dao *Dao
}

var smOnce sync.Once
var sm sessionManager

// SessionManager exportable singleton
func SessionManager(dao *Dao) *sessionManager {
	smOnce.Do(func() {
		sm = sessionManager{dao}
	})
	return &sm
}

// SessionManagerActions exportable
func (m sessionManager) Actions() (ac Actions, err error) {
	return
}
