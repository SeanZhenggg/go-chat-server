package errortool

import (
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

const (
	DBGroupCode int = 3

	// pg native error code
	UniqueViolationCode string = pgerrcode.UniqueViolation
)

func ProvideDBError(groups IGroupRepo, codes ICodeRepo) interface{} {
	group := Define.GenGroup(DBGroupCode)

	return &dbError{
		NoRow:           group.GenError(1, "No Rows Returned"),
		UniqueViolation: group.GenError(2, "Duplicate Key Value Violation"),
	}
}

var (
	dbErrorCodeMap = map[string]error{
		UniqueViolationCode: DbErr.UniqueViolation,
	}
)

type dbError struct {
	NoRow           error
	UniqueViolation error
}

func ParseDBError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return DbErr.NoRow
	}

	var pgError *pgconn.PgError
	ok := errors.As(err, &pgError)
	if !ok {
		return err
	}

	if parseErr, ok := dbErrorCodeMap[pgError.Code]; ok {
		return parseErr
	}

	return err
}
