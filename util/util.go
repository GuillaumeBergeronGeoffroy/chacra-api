/*
	Package util ...
*/
package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

// ReadUserIP exportable
func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

// CreateUUID exportable
func CreateUUID() string {
	id := uuid.New()
	return id.String()
}

// Read exportable
func Read(w http.ResponseWriter, r *http.Request) (reqBody []byte) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

// Write exportable
func Write(w http.ResponseWriter, r *http.Request, resBody []byte) {
	if resBody != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resBody)
	}
}

// Request exportable
func Request(route string, reqBody []byte) (resStatusCode int, resBody []byte, err error) {
	resp, err := http.Post(route, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	resStatusCode = resp.StatusCode
	resBody, err = ioutil.ReadAll(resp.Body)
	return
}

// ComposeResponse exportable
func ComposeResponse(w http.ResponseWriter, resp map[string]string) (resBody []byte) {
	resBody, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}
