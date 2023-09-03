package errortool

import "log"

type ICodeRepo interface {
	Add(code int, err *CustomError)
	Get(code int) (*CustomError, bool)
}

func ProvideCodeRepo() ICodeRepo {
	return &codeRepo{}
}

type codeRepo struct {
	store map[int]*CustomError
}

func (c *codeRepo) Add(code int, err *CustomError) {
	if _, ok := c.store[code]; ok {
		log.Panicf("error code duplicate definition, code: %d", code)
	}

	c.store[code] = err
}

func (c *codeRepo) Get(code int) (*CustomError, bool) {
	val, ok := c.store[code]
	return val, ok
}
