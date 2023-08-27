package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() {
	_, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        "host=localhost user=postgres password=postgrespw dbname=postgres port=55000",
	}))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("connection success!!!")
}

func main() {
	Init()
}
