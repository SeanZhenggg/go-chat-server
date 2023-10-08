package bo

import "time"

type GetUserCond struct {
	Account string
}

type UserLoginCond struct {
	Account  string
	Password string
}

type UserRegCond struct {
	Account  string
	Password string
	Nickname string
}

type UserInfo struct {
	Id          uint
	Account     string
	Password    string
	Nickname    string
	Birthdate   time.Time
	Gender      int
	Country     string
	Address     string
	RegionCode  string
	PhoneNumber string
	CreateAt    time.Time
	UpdateAt    time.Time
}

type UserLoginResp struct {
	Token string
}

type UserValidateCond struct {
	Token string
}

type UpdateUserInfoCond struct {
	Id          uint
	Password    string
	Nickname    string
	Birthdate   time.Time
	Gender      int
	Country     string
	Address     string
	RegionCode  string
	PhoneNumber string
}
