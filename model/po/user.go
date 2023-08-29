package po

import (
	"time"
)

type User struct {
	Id       uint
	Account  string
	Password string
	Nickname string
	CreateAt time.Time
	UpdateAt time.Time
}

func (User) TableName() string {
	return "users"
}

type UserCond struct {
	Account string
}

type UserRegData struct {
	Account  string
	Password string
	Nickname string
}
