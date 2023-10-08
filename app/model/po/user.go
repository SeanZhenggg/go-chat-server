package po

import (
	"time"
)

type User struct {
	BaseId
	Account     string    `gorm:"column:account"`
	Password    string    `gorm:"column:password"`
	Birthdate   time.Time `gorm:"column:birthdate;default:null"`
	Gender      int       `gorm:"column:gender;default:1"`
	Country     string    `gorm:"column:country;default:null"`
	Address     string    `gorm:"column:address;default:null"`
	RegionCode  string    `gorm:"column:region_code;default:null"`
	PhoneNumber string    `gorm:"column:phone_number;default:null"`
	Nickname    string    `gorm:"column:nickname;default:null"`
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

type UpdateUserInfoIdCond struct {
	BaseId
}

type UpdateUserInfoCond struct {
	Password    *string    `gorm:"column:password"`
	Birthdate   *time.Time `gorm:"column:birthdate"`
	Gender      *int       `gorm:"column:gender"`
	Country     *string    `gorm:"column:country"`
	Address     *string    `gorm:"column:address"`
	RegionCode  *string    `gorm:"column:region_code"`
	PhoneNumber *string    `gorm:"column:phone_number"`
	Nickname    *string    `gorm:"column:nickname"`
}
