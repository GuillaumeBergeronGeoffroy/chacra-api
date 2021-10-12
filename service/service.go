/*
	Package service ...
*/
package service

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/go-redis/redis/v8"
)

// ServiceConfig exportable
type ServiceConfig struct {
	name       string
	host       string
	user       string
	password   string
	mysqldb    bool
	redisStore bool
}

// Action exportable
type Action func(w http.ResponseWriter, r *http.Request)

// Actions exportable
type Actions map[string]Action

type dao struct {
	db          *sql.DB
	ctx         context.Context
	redisClient *redis.Client
}

var daoIns = dao{}

// InitService exportable
func InitService(sC ServiceConfig, gateway map[string]string) (ac Actions, err error) {
	// Mysql connection pool init
	if sC.mysqldb && daoIns.db == nil {
		daoIns.db, err = sql.Open("mysql", sC.user+"@"+sC.password)
		if err != nil {
			return
		}
	}
	// Redis connection pool init
	if sC.redisStore && daoIns.redisClient == nil {
		daoIns.ctx = context.TODO()
		daoIns.redisClient = redis.NewClient(&redis.Options{
			Addr:     sC.host,
			Password: sC.password,
			DB:       0,
		})
		if err = daoIns.redisClient.Ping(daoIns.ctx).Err(); err != nil {
			return
		}
	}
	switch sC.name {
	case "ClientPortal":
		ac, err = ClientPortal(daoIns.db, gateway).Actions()
	case "ContentManager":
		ac, err = ContentManager(daoIns.db, gateway).Actions()
	case "EmployeePortal":
		ac, err = EmployeePortal(daoIns.db, gateway).Actions()
	case "LogManager":
		ac, err = LogManager(daoIns.db, gateway).Actions()
	case "ProducerPortal":
		ac, err = ProducerPortal(daoIns.db, gateway).Actions()
	case "SessionManager":
		ac, err = SessionManager(daoIns.redisClient, gateway).Actions()
	case "TransactionManager":
		ac, err = TransactionManager(daoIns.db, gateway).Actions()
	}
	return
}
