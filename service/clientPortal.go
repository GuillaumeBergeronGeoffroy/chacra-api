package service

import (
	"encoding/json"
	"net/http"
	"sync"

	u "github.com/GuillaumeBergeronGeoffroy/chacra-api/util"
	"golang.org/x/crypto/bcrypt"
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
			s := &SubscribeRequest{}
			reqBody := u.Read(w, r)
			err = json.Unmarshal([]byte(reqBody), s)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// check if user exists
			// Validate email and password
			var insertStmt string
			if s.UserPassword == "" {
				insertStmt = "INSERT INTO User (UserEmail, UserStatus) VALUES ('" + s.UserEmail + "', 0)"
			} else {
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s.UserPassword), bcrypt.DefaultCost)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				insertStmt = "INSERT INTO User (UserEmail, UserPassword, UserStatus) VALUES ('" + s.UserEmail + "','" + string(hashedPassword) + "', 1)"
			}
			stmt := []string{
				insertStmt,
			}
			err = ExecuteStatements(m.Dao.DB, stmt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
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
