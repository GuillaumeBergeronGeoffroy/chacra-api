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
	})
	return &cp
}

// ClientPortalActions exportable
func (m clientPortal) Actions() (ac Actions, err error) {
	return
}
