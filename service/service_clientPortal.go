package service

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

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
			if err = json.Unmarshal([]byte(reqBody), user); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			user.UserCreatedAt = time.Now().Format("2006-01-02 15:04:05")
			if err = SaveModel(user, m.Dao.DB); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(ComposeResponse(w, map[string]string{
				"message": "Votre place est réservé!",
				"success": "true",
			}))
		},
	}
	return
}

var cpInitSql = []string{
	`CREATE TABLE User (
		UserId INT NOT NULL AUTO_INCREMENT,
		UserEmail VARCHAR(255) NOT NULL,
		UserPassword VARCHAR(55), 
		UserName VARCHAR(255), 
		UserCreatedAt DATETIME default current_timestamp,
		UserStatus TINYINT DEFAULT 0,
		PRIMARY KEY (UserId),
		CONSTRAINT uidx_User_UserEmail UNIQUE (UserEmail)
	);`,
}
