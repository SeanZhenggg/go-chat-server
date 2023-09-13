package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbDSN    = "host=localhost user=postgres password=postgrespw dbname=postgres port=55000"
	dbDriver = "pgx"
)

type IPostgresDB interface {
	Session() *gorm.DB
}

func ProvidePostgresDB() IPostgresDB {
	newDB := &postgresDB{}
	newDB.DB = newDB.postgresDBConnect()

	return newDB
}

type postgresDB struct {
	DB *gorm.DB
}

func (p *postgresDB) Session() *gorm.DB {
	return p.DB.Session(&gorm.Session{SkipDefaultTransaction: true})
}

func (p *postgresDB) postgresDBConnect() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: dbDriver,
		DSN:        dbDSN,
	}))

	if err != nil {
		panic(fmt.Sprintf("open mysql error and error message is '%v'", err))
	}

	return db
}
