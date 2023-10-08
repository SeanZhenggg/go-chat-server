package po

import (
	"time"
)

type User struct {
	BaseId
	Account     string     `gorm:"column:account"`
	Password    string     `gorm:"column:password"`
	Birthdate   *time.Time `gorm:"column:birthdate"`
	Gender      int        `gorm:"column:gender"`
	Country     string     `gorm:"column:country"`
	Address     string     `gorm:"column:address"`
	RegionCode  string     `gorm:"column:region_code"`
	PhoneNumber string     `gorm:"column:phone_number"`
	Nickname    string     `gorm:"column:nickname"`
	BaseTimeColumns
}

func (User) TableName() string {
	return "users"
}

type GetUserCond struct {
	Account string
}

type UserRegCond struct {
	Account  string
	Password string
	Nickname string
}

type UserLoginCond struct {
	Account  string
	Password string
}

type UpdateUserInfoCond struct {
	BaseId
	Password    string    `gorm:"column:password"`
	Birthdate   time.Time `gorm:"column:birthdate"`
	Gender      int       `gorm:"column:gender"`
	Country     string    `gorm:"column:country"`
	Address     string    `gorm:"column:address"`
	RegionCode  string    `gorm:"column:region_code"`
	PhoneNumber string    `gorm:"column:phone_number"`
	Nickname    string    `gorm:"column:nickname"`
}
