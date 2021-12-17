package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ComposeResponse exportable
func ComposeResponse(w http.ResponseWriter, resp map[string]string) (resBody []byte) {
	resBody, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}


// Read exportable
func Read(w http.ResponseWriter, r *http.Request) (reqBody []byte) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

// Request exportable
func Request(netClient *http.Client, route string, reqBody []byte) (resStatusCode int, resBody []byte, err error) {
	resp, err := netClient.Post(route, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	resStatusCode = resp.StatusCode
	resBody, err = ioutil.ReadAll(resp.Body)
	return
}

// WriteJSON exportable
func WriteJSON(w http.ResponseWriter, r *http.Request, resBody []byte) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resBody)
}

