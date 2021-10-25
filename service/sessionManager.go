package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"

	u "github.com/GuillaumeBergeronGeoffroy/chacra-api/util"
)

type sessionManager struct {
	Dao *Dao
}

var smOnce sync.Once
var sm sessionManager

// SessionManager exportable singleton
func SessionManager(dao *Dao) *sessionManager {
	smOnce.Do(func() {
		sm = sessionManager{dao}
		dao.RateLimiter = &RateLimiter{RateMap: map[string]RateEntry{}, RateDelay: 120, BanDelay: 60 * 60, Limit: 12}
	})
	return &sm
}

type SubscribeRequest struct {
	UserEmail        string `json:"userEmail,omitempty"`
	ProducerEmail    string `json:"producerEmail,omitempty"`
	UserPassword     string `json:"userPassword,omitempty"`
	ProducerPassword string `json:"producerPassword,omitempty"`
}

func subscribe(reqBody []byte, m sessionManager) (resBody []byte, err error) {
	s := &SubscribeRequest{}
	err = json.Unmarshal([]byte(reqBody), s)
	if err != nil {
		return
	}
	var gateway string
	var endRoute string
	if s.ProducerEmail != "" && s.ProducerPassword != "" {
		gateway = "ProducerClient"
		endRoute = "/createProducer"
	} else if s.UserEmail != "" && s.UserPassword != "" {
		gateway = "UserClient"
		endRoute = "/createUser"
	} else {
		err = errors.New("invalid submission format")
		return
	}
	resStatus, resBody, err := u.Request(m.Dao.Gateway[gateway]+endRoute, reqBody)
	if resStatus < 200 || resStatus > 299 {
		err = errors.New("something went wrong")
	}
	return
}

// SessionManagerActions exportable
func (m sessionManager) Actions() (ac Actions, err error) {
	ac = map[string]Action{
		"subscribe": func(w http.ResponseWriter, r *http.Request) {
			err = EvalRateLimit(r, m.Dao.RateLimiter)
			if err != nil {
				u.Write(w, r, u.ComposeResponse(w, map[string]string{
					"message": err.Error(),
					"success": "false",
				}))
				return
			}
			reqBody := u.Read(w, r)
			resBody, err := subscribe(reqBody, m)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			u.Write(w, r, resBody)
			return
		},
		"authentify": func(w http.ResponseWriter, r *http.Request) {
			return
		},
		"authorize": func(w http.ResponseWriter, r *http.Request) {
			return
		},
	}
	return
}
