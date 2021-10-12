package service

import (
	"database/sql"
	"sync"
)

type contentManager struct {
	db      *sql.DB
	gateway map[string]string
}

var cmOnce sync.Once
var cm contentManager

// ContentManager singleton exportable
func ContentManager(db *sql.DB, gateway map[string]string) contentManager {
	cmOnce.Do(func() {
		cm = contentManager{db, gateway}
	})
	return cm
}

// ContentManagerActions exportable
func (m contentManager) Actions() (ac Actions, err error) {
	return
}
