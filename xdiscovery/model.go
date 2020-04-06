// This is free and unencumbered software released into the public domain.

// Anyone is free to copy, modify, publish, use, compile, sell, or
// distribute this software, either in source code form or as a compiled
// binary, for any purpose, commercial or non-commercial, and by any
// means.

// In jurisdictions that recognize copyright laws, the author or authors
// of this software dedicate any and all copyright interest in the
// software to the public domain. We make this dedication for the benefit
// of the public at large and to the detriment of our heirs and
// successors. We intend this dedication to be an overt act of
// relinquishment in perpetuity of all present and future rights to this
// software under copyright law.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

// For more information, please refer to <https://unlicense.org>

// Package discovery implements a service registry for tracking the location of
// distributed microservices.
package xdiscovery

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"
)

// Service holds information about a service as well as the last time the
// service was renewed.
type Service struct {
	Name  string    `json:"name"`
	Host  string    `json:"host"`
	Added time.Time `json:"added"`
}

// Registry holds host names for services by name.
type Registry interface {
	Add(service Service)              // Add adds or updates a service to this registry.
	Remove(service Service)           // Remove removes a service from this registry.
	Get(name string) (Service, error) // Get gets the specified service.
	List(name string) []Service       // List gets all services filtered by name.
	SetTimeout(timeout time.Duration) // SetTimeout updates the timeout duration.
	SetKeep(timeout time.Duration)    // SetKeep updates the keep duration.
}

// Authenticator defines how to handle http authentication.
type Authenticator func(token string) bool

// NullAuthenticator the authenticator that always returns true.
func NullAuthenticator(token string) bool { return true }

// NewBasicAuthenticator returns an authenticator that performs basic auth with
// the supplied username and password.
func NewBasicAuthenticator(username, password string) Authenticator {
	cred := []byte(fmt.Sprintf("%s:%s", username, password))
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(cred)
	return func(token string) bool {
		return token == string(cred)
	}
}
