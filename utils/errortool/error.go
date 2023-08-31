package errortool

import "errors"

func UnwrapErrors(err error) (error, bool) {
	newError := err
	for {
		if tmp := errors.Unwrap(newError); tmp != nil {
			newError = tmp
		} else {
			break
		}
	}

	if parsed, ok := newError.(error); ok {
		return parsed, true
	} else {
		return nil, false
	}
}
