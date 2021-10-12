package service

import (
	"database/sql"
	"sync"
)

type clientPortal struct {
	db      *sql.DB
	gateway map[string]string
}

var cpOnce sync.Once
var cp clientPortal

// ClientPortal singleton exportable
func ClientPortal(db *sql.DB, gateway map[string]string) clientPortal {
	cpOnce.Do(func() {
		cp = clientPortal{db, gateway}
	})
	return cp
}

// ClientPortalActions exportable
func (m clientPortal) Actions() (ac Actions, err error) {
	return
}
