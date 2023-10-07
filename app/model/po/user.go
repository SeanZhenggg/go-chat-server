package po

import (
	"time"
)

type User struct {
	Id          uint      `gorm:"id"`
	Account     string    `gorm:"account"`
	Password    string    `gorm:"password"`
	Birthdate   time.Time `gorm:"birthdate"`
	Gender      int       `gorm:"gender"`
	Country     string    `gorm:"country"`
	Address     string    `gorm:"address"`
	RegionCode  string    `gorm:"region_code"`
	PhoneNumber string    `gorm:"phone_number"`
	Nickname    string    `gorm:"nickname"`
	CreateAt    time.Time `gorm:"column:create_at;autoCreateTime"`
	UpdateAt    time.Time `gorm:"column:update_at;autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}

type UserCond struct {
	Account string
}

type UserRegData struct {
	Account  string `gorm:"account"`
	Password string `gorm:"password"`
	Nickname string `gorm:"nickname"`
}

type UserLoginData struct {
	Account  string `gorm:"password"`
	Password string `gorm:"nickname"`
}
