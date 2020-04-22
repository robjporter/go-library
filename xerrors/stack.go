package xerrors

func FirstStackTrace(err error) StackTrace {
	var st StackTrace
	if err == nil {
		return nil
	}
	for {
		if s, ok := err.(StackTracer); ok {
			st = s
		}
		if c, ok := err.(Causer); ok {
			if prev = c.Cause(); prev != nil {
				err = prev
				continue
			}
		}
		break
	}
	return st
}

func HasStackTrace(err error) bool {
	if err == nil {
		return false
	}
	for {
		if _, ok := err.(StackTracer); ok {
			return true
		}
		if c, ok := err.(Causer); ok {
			if prev = c.Cause(); prev != nil {
				err = prev
				continue
			}
		}
		break
	}
	return false
}
