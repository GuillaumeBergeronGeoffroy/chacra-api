/*
	Package util
		- general funcs by (descriptive) name (alphabetical)
*/
package util

import (
	"net/http"

	"github.com/google/uuid"
)

// CreateUUID exportable
func CreateUUID() string {
	id := uuid.New()
	return id.String()
}

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
