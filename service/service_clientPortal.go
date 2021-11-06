package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	model "github.com/GuillaumeBergeronGeoffroy/chacra-api/model"
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
		ExecuteStatements(cp.Dao.DB, cpInitSql)
	})
	return &cp
}

// ClientPortalActions exportable
func (m clientPortal) Actions() (ac Actions, err error) {
	ac = map[string]Action{
		"createUser": func(w http.ResponseWriter, r *http.Request) {
			user := &model.User{}
			reqBody := Read(w, r)
			err = json.Unmarshal([]byte(reqBody), user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Println(user)
			err = SaveModel(user, m.Dao.DB)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		},
	}
	return
}

var cpInitSql = []string{
	`CREATE TABLE User (
		UserId INT NOT NULL AUTO_INCREMENT,
		UserEmail VARCHAR(255) NOT NULL,
		UserPassword VARCHAR(255), 
		UserCreatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UserStatus TINYINT DEFAULT 0,
		PRIMARY KEY (UserId),
		CONSTRAINT uidx_User_UserEmail UNIQUE (UserEmail)
	);`,
}
