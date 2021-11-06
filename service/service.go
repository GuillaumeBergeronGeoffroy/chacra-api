/*
	Package service
		- service
			built to be run independantly, defined by config, need gateway map to other services
				structure
					parts by hierarchy
					singleton init
					Actions by name (alphabelical)
					functions used by the Actions in order of appearance
					service init sql stmt
		- helper
			grouped functions
				structure
					by name (alphabelical)
		- tools
			object groups logic that solves a problem
				structure
					parts by hierarchy
					funcs by name (alphabelical)
*/
package service

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
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

// Dao exportable
type Dao struct {
	DB          *sql.DB
	Ctx         context.Context
	RedisClient *redis.Client
	Gateway     map[string]string
	RateLimiter *RateLimiter
	HttpClient  *http.Client
}

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

// InitService exportable
func InitService(sc ServiceConfig, gateway map[string]string) (ac Actions, err error) {
	var daoIns = Dao{}
	// Set Gateway
	daoIns.Gateway = gateway
	// The HTTPClient
	daoIns.HttpClient = netClient
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
