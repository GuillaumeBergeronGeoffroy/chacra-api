package service

import "github.com/go-redis/redis/v8"

// ProducerPortal exportable
type SessionManager struct {
	redisClient *redis.Client
	links       map[string]string
}

// SessionManager exportable
func (m SessionManager) Actions() (ac Actions, err error) {
	return
}
