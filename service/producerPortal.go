package service

import (
	"database/sql"
	"sync"
)

type producerPortal struct {
	db      *sql.DB
	gateway map[string]string
}

var ppOnce sync.Once
var pp producerPortal

// ProducerPortal exportable singleton
func ProducerPortal(db *sql.DB, gateway map[string]string) producerPortal {
	ppOnce.Do(func() {
		pp = producerPortal{db, gateway}
	})
	return pp
}

// ProducerPortalActions exportable
func (m producerPortal) Actions() (ac Actions, err error) {
	return
}
