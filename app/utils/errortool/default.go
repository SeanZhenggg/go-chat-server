package errortool

import (
	"log"
)

const (
	ErrorCodeMin   = 1
	ErrorCodeMax   = 999
	ErrorCodeRange = 1000
)

func ProvideDefine() *define {
	return &define{
		codes:  ProvideCodeRepo(),
		groups: ProvideGroupRepo(),
	}
}

type define struct {
	codes  ICodeRepo
	groups IGroupRepo
}

func genGroupCode(groupBase int) int {
	return groupBase * ErrorCodeRange
}

func (d *define) GenGroup(base int) *errorGroup {
	baseGroupCode := genGroupCode(base)
	d.groups.Add(baseGroupCode)

	return &errorGroup{
		codes:  d.codes,
		groups: d.groups,
		group:  baseGroupCode,
	}
}

func (d *define) Plugin(f func(groups IGroupRepo, codes ICodeRepo) interface{}) interface{} {
	return f(d.groups, d.codes)
}

type errorGroup struct {
	codes  ICodeRepo
	groups IGroupRepo
	group  int
}

func (e *errorGroup) GenError(code int, message string) error {
	if code > ErrorCodeMax || code < ErrorCodeMin {
		log.Panicf("errorGroup error: code less than %v or code greater than %v, code: %d", ErrorCodeMin, ErrorCodeMax, code)
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

func (e *errorGroup) makeCode(code int) int {
	return e.groups.Get(e.group) + code
}
