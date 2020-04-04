package xenv

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"
	"strconv"
	"time"
)

type Env struct {
	vars map[string]string
}

func New() *Env {
	return &Env{vars: make(map[string]string)}
}

func (e *Env) Size() int {
	return len(e.vars)
}

func (e *Env) AddString(key, value string) {
	if os.Getenv(key) == "" {
		e.vars[key] = value
		return
	}
	e.vars[key] = os.Getenv(key)
}

func (e *Env) GetAll() map[string]string {
	return e.vars
}

func (e *Env) SetString(key, value string) error {
	e.vars[key] = value
	return os.Setenv(key, value)
}

func (e *Env) SetBool(key string, value bool) error {
	val := ""
	if(value) {
		val = "true"
	} else {
		val = "false"
	}
	e.vars[key] = val
	return os.Setenv(key, val)
}

func (e *Env) SetInt(key string, value int) error {
	val := strconv.Itoa(value)
	e.vars[key] = val
	return os.Setenv(key, val)
}

func (e *Env) GetString(key string) string {
	if e.vars[key] == "" {
		v := os.Getenv(key)
		if v != "" {
			e.vars[key] = v
			return e.vars[key]
		}
		return ""
	}
	return e.vars[key]
}

func (e *Env) GetAsBool(key string) bool {
	val := strings.ToLower(strings.TrimSpace(e.Get(key)))
	switch val {
	case "0", "no", "false":
		return false
	default:
		if val != "" {
			return true
		}
		return false
	}
}

func (e *Env) GetAsInt(key string) int {
	i, _ := strconv.Atoi(strings.ToLower(strings.TrimSpace(e.Get(key))))
	return i
}

func (e *Env) GetUsername() string {
	u, e := user.Current()
	if e != nil {
		return ""
	}
	return u.Username
}

func (e *Env) GOPATHBIN() string {
	return os.Getenv("GOPATH") + PathSeparator() + "bin"
}

func (e *Env) PathSeparator() string {
	return fmt.Sprintf("%c", os.PathSeparator)
}

func (e *Env) ListSeparator() string {
	return fmt.Sprintf("%c", os.PathListSeparator)
}

func (e *Env) IsCompiled() bool {
	if strings.HasPrefix(os.Args[0], "/var/folders/") ||
		strings.HasPrefix(os.Args[0], "/tmp/go-build") ||
		strings.Contains(os.Args[0], "\\AppData\\Local\\Temp\\") {
		return false
	}
	return true
}

func (e *Env) BuildDebug() bool {
	return debug
}

func (e *Env) CheckArchitecture() bool {
	switch runtime.GOARCH {
	case "386", "amd64":
		return true
	case "arm64", "arm", "ppc64", "ppc64le", "mips", "mipsle":
		fmt.Println("This is an untested architecture: %q; proceed with caution!", runtime.GOARCH)
		return false
	default:
		fmt.Printf("Unknown goarch %q; proceed with caution!", runtime.GOARCH)
		return false
	}
}

func (e *Env) BuildStamp() int64 {
	if s, _ := strconv.ParseInt(os.Getenv("SOURCE_DATE_EPOCH"), 10, 64); s > 0 {
		return s
	}
	bs, err := runError("git", "show", "-s", "--format=%ct")
	if err != nil {
		return time.Now().Unix()
	}
	s, _ := strconv.ParseInt(string(bs), 10, 64)
	return s
}


func (e *Env) Compiler() string {
	return runtime.Compiler
}

func (e *Env) GOARCH() string {
	return runtime.GOARCH
}

func (e *Env) GOOS() string {
	return runtime.GOOS
}

func (e *Env) GOROOT() string {
	return runtime.GOROOT()
}

func (e *Env) GOVER() string {
	return runtime.Version()
}

func (e *Env) NumCPU() int {
	return runtime.NumCPU()
}

func (e *Env) GOPATH() string {
	return os.Getenv("GOPATH")
}

func (e *Env) GetFormattedTime() string {
	return Now("Monday, 2 Jan 2006")
}

func (e *Env) Now(layout string) string {
	return time.Now().Format(layout)
}