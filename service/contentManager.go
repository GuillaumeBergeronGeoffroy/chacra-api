package service

import (
	"sync"
)

type contentManager struct {
	Dao *Dao
}

var cmOnce sync.Once
var cm contentManager

// ContentManager singleton exportable
func ContentManager(dao *Dao) *contentManager {
	cmOnce.Do(func() {
		cm = contentManager{dao}
	})
	return &cm
}

// ContentManagerActions exportable
func (m contentManager) Actions() (ac Actions, err error) {
	return
}
