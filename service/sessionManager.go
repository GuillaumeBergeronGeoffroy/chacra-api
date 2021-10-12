package service

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

type sessionManager struct {
	redisClient *redis.Client
	gateway     map[string]string
}

var smOnce sync.Once
var sm sessionManager

// SessionManager exportable singleton
func SessionManager(redisClient *redis.Client, gateway map[string]string) sessionManager {
	smOnce.Do(func() {
		sm = sessionManager{redisClient, gateway}
	})
	return sm
}

// SessionManagerActions exportable
func (m sessionManager) Actions() (ac Actions, err error) {
	return
}
