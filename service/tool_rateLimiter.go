package service

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	util "github.com/GuillaumeBergeronGeoffroy/chacra-api/util"
)

// RateEntry exportable
type RateEntry struct {
	Count      int
	LastAction time.Time
	Ban        bool
}

// RateLimiter exportable
type RateLimiter struct {
	mu        sync.Mutex
	RateMap   map[string]RateEntry
	RateDelay float64
	BanDelay  float64
	Limit     int
}

// EvalRateLimit exportable
func EvalRateLimit(r *http.Request, rateLimiter *RateLimiter) (err error) {
	fmt.Println(rateLimiter.RateMap)
	ip := util.ReadUserIP(r)
	if ip == "" {
		return
	}
	rateLimiter.mu.Lock()
	defer rateLimiter.mu.Unlock()
	if rateEntry, ok := rateLimiter.RateMap[ip]; ok {
		if rateEntry.Ban == true {
			if time.Now().Sub(rateEntry.LastAction).Seconds() > rateLimiter.BanDelay {
				rateLimiter.RateMap[ip] = RateEntry{Count: 1, LastAction: time.Now(), Ban: false}
				err = errors.New("gotta go fast")
			}
		} else {
			if time.Now().Sub(rateEntry.LastAction).Seconds() > rateLimiter.RateDelay {
				rateLimiter.RateMap[ip] = RateEntry{Count: 1, LastAction: time.Now(), Ban: false}
				err = errors.New("gotta go fast")
			} else {
				if rateLimiter.RateMap[ip].Count+1 > rateLimiter.Limit {
					rateLimiter.RateMap[ip] = RateEntry{Count: rateLimiter.RateMap[ip].Count + 1, LastAction: time.Now(), Ban: true}
				} else {
					rateLimiter.RateMap[ip] = RateEntry{Count: rateLimiter.RateMap[ip].Count + 1, LastAction: rateLimiter.RateMap[ip].LastAction, Ban: false}
				}
			}
		}
	} else {
		rateLimiter.RateMap[ip] = RateEntry{Count: 0, LastAction: time.Now(), Ban: false}
	}
	return
}
