package errortool

import (
	"log"
)

const (
	errCodeMax = 999
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
	return groupBase * 1000
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
	if code > errCodeMax {
		log.Panicf("errorGroup error: code max than 999, code: %d", code)
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
