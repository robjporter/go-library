package xerrors

func (rc RootCause) Error() string {
	return rc.s
}

func (rc RootCause) StackTrace() StackTrace {
	return rc.st
}
