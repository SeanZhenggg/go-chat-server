package errortool

import (
	"log"
)

const (
	DefaultErrorCodeMin   = 1
	DefaultErrorCodeMax   = 999
	DefaultErrorCodeRange = 1000
)

func ProvideDefine() *Define {
	return &Define{
		codes: &codeRepo{
			store: make(map[int]*CustomError, 256),
		},
		groups: &groupRepo{
			store: make(map[int]struct{}, 256),
		},
		errCodeMin:   DefaultErrorCodeMin,
		errCodeMax:   DefaultErrorCodeMax,
		errCodeRange: DefaultErrorCodeRange,
	}
}

func (d *Define) genGroupCode(groupBase int) int {
	return groupBase * d.errCodeRange
}

type Define struct {
	codes        ICodeRepo
	groups       IGroupRepo
	errCodeMin   int
	errCodeMax   int
	errCodeRange int
}

type ErrCodeOptions struct {
	Min   int
	Max   int
	Range int
}

func (d *Define) GenErrorGroup(base int) *ErrorGroup {
	baseGroupCode := d.genGroupCode(base)
	d.groups.Add(baseGroupCode)

	return &ErrorGroup{
		codes:  d.codes,
		groups: d.groups,
		group:  baseGroupCode,
		define: d,
	}
}

func (d *Define) SetErrCodeOptions(opt ErrCodeOptions) *Define {
	if opt.Range < opt.Min || opt.Range < opt.Max {
		log.Fatalf(
			"errortool Define.SetErrCodeOptions error : range [%d] is less than min [%d] or max [%d]",
			opt.Range,
			opt.Min,
			opt.Max,
		)
	}

	return d
}

type ErrorGroup struct {
	codes  ICodeRepo
	groups IGroupRepo
	group  int
	define *Define
}

func (e *ErrorGroup) GenError(code int, message string) error {
	if code < e.define.errCodeMin || code > e.define.errCodeMax {
		log.Panicf(
			"errortool ErrorGroup.GenError error: code [%d] less than %v or greater than %v",
			code,
			e.define.errCodeMin,
			e.define.errCodeMax,
		)
	}

	errCode := e.makeCode(code)
	err := &CustomError{
		code:      errCode,
		baseCode:  code,
		groupCode: e.group,
		message:   message,
	}
	e.codes.Add(errCode, err)

	return err
}

func (e *ErrorGroup) makeCode(code int) int {
	return e.groups.Get(e.group) + code
}
