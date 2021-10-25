/*
	Package service ...
*/
package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"

	u "github.com/GuillaumeBergeronGeoffroy/chacra-api/util"
)

// ServiceConfig exportable
type ServiceConfig struct {
	Name       string
	Host       string
	Port       string
	User       string
	Password   string
	Mysqldb    bool
	RedisStore bool
}

// Service Exportable
type Service interface {
	Actions(ac Actions, err error)
}

// Action exportable
type Action func(w http.ResponseWriter, r *http.Request)

// Actions exportable
type Actions map[string]Action

// RateEntry exportable
type RateEntry struct {
	Count      int
	LastAction time.Time
	Ban        bool
}

// RateLimiter exportable
type RateLimiter struct {
	RateMap   map[string]RateEntry
	RateDelay float64
	BanDelay  float64
	Limit     int
}

// Dao exportable
type Dao struct {
	DB          *sql.DB
	Ctx         context.Context
	RedisClient *redis.Client
	Gateway     map[string]string
	RateLimiter *RateLimiter
}

// EvalRateLimit exportable
func EvalRateLimit(r *http.Request, rateLimiter *RateLimiter) (err error) {
	ip := u.ReadUserIP(r)
	if ip == "" {
		return
	}
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

// ExecuteStatements exportable
func ExecuteStatements(db *sql.DB, stmts []string) (err error) {
	for _, stmt := range stmts {
		ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelfunc()
		_, err = db.ExecContext(ctx, stmt)
		if err != nil {
			return
		}
	}
	return
}

// InitService exportable
func InitService(sc ServiceConfig, gateway map[string]string) (ac Actions, err error) {
	var daoIns = Dao{}
	// Set Gateway
	daoIns.Gateway = gateway
	// Mysql connection pool init
	if sc.Mysqldb {
		daoIns.DB, err = sql.Open("mysql", sc.User+":"+sc.Password+"@tcp("+sc.Host+":"+sc.Port+")/chacra")
		if err != nil {
			return
		}
		err = daoIns.DB.Ping()
		if err != nil {
			return
		}
		defer daoIns.DB.Close()
	}
	// Redis connection pool init
	if sc.RedisStore {
		daoIns.Ctx = context.TODO()
		daoIns.RedisClient = redis.NewClient(&redis.Options{
			Addr:     sc.Host + ":" + sc.Port,
			Password: sc.Password,
			DB:       0,
		})
		if err = daoIns.RedisClient.Ping(daoIns.Ctx).Err(); err != nil {
			return
		}
		defer daoIns.RedisClient.Close()
	}
	switch sc.Name {
	case "ClientPortal":
		ac, err = ClientPortal(&daoIns).Actions()
	case "ContentManager":
		ac, err = ContentManager(&daoIns).Actions()
	case "EmployeePortal":
		ac, err = EmployeePortal(&daoIns).Actions()
	case "LogManager":
		ac, err = LogManager(&daoIns).Actions()
	case "ProducerPortal":
		ac, err = ProducerPortal(&daoIns).Actions()
	case "SessionManager":
		ac, err = SessionManager(&daoIns).Actions()
	case "TransactionManager":
		ac, err = TransactionManager(&daoIns).Actions()
	}
	return
}
