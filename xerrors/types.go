package xerrors

import "runtime"

type Causer interface {
	Cause() error
}

type Domino struct {
	cause  error
	effect error
	s      string
	st     StackTrace
}

type Frame *runtime.Frame

type RootCause struct {
	s  string
	st StackTrace
}

type StackTrace []Frame

type StackTracer interface {
	StackTrace() StackTrace
}
