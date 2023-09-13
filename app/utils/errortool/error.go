package errortool

import (
	"errors"
	"fmt"
)

func ParseError(err error) (*CustomError, bool) {
	newError := err
	for {
		if tmp := errors.Unwrap(newError); tmp != nil {
			newError = tmp
		} else {
			break
		}
	}

	if parsed, ok := newError.(*CustomError); ok {
		return parsed, true
	} else {
		return nil, false
	}
}

type CustomError struct {
	code      int
	baseCode  int
	groupCode int
	message   string
}

func (e *CustomError) Error() string {
	return fmt.Sprint(e.code) + ":" + e.message
}

func (e *CustomError) GetCode() int {
	return e.code
}

func (e *CustomError) GetMessage() string {
	return e.message
}
