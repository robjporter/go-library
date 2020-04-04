package xdebug

import (
	"os"
	"strings"
)

var _GODEBUG_ITEMS_ = make(map[string]string)

func init() {
	tmp := os.Environ()
	for i := 0; i < len(tmp); i++ {
		item := strings.Split(tmp[i], "=")
		_GODEBUG_ITEMS_[item[0]] = item[1]
	}
}

func GODEBUGDATA() map[string]string {
	return _GODEBUG_ITEMS_
}

// Get setting value in GODEBUG variable by name.
// If setting not exists, returns defaultValue.
// See GODEBUG environment variable description in http://golang.org/pkg/runtime/
func GODEBUG(name, defaultValue string) string {
	if v, ok := _GODEBUG_ITEMS_[name]; ok {
		return v
	}
	return defaultValue
}