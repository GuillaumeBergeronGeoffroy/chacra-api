package service

import (
	"sync"
)

type producerPortal struct {
	Dao *Dao
}

var ppOnce sync.Once
var pp producerPortal

// ProducerPortal exportable singleton
func ProducerPortal(dao *Dao) *producerPortal {
	ppOnce.Do(func() {
		pp = producerPortal{dao}
	})
	return &pp
}

// ProducerPortalActions exportable
func (m producerPortal) Actions() (ac Actions, err error) {
	return
}
