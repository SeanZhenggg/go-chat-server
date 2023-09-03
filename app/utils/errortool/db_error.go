package errortool

import (
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	groupCodeDB int = 10

	// pg native error code
	UniqueViolationCode string = pgerrcode.UniqueViolation
)

func ProvideDBError(groups IGroupRepo, codes ICodeRepo) interface{} {
	group := Define.GenGroup(groupCodeDB)

	return &dbError{
		UniqueViolation: group.GenError(1, "Duplicate Key Value Violation"),
	}
}

var (
	dbErrorCodeMap = map[string]error{
		UniqueViolationCode: DbErr.UniqueViolation,
	}
)

type dbError struct {
	UniqueViolation error
}

func ParseDBError(err error) error {
	pgError, ok := err.(*pgconn.PgError)
	if !ok {
		return err
	}

	if parseErr, ok := dbErrorCodeMap[pgError.Code]; ok {
		return parseErr
	}

	return err
}
