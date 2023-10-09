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
	CountryCode string
	Address     string
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

type UpdateUserInfoIdCond struct {
	Id uint
}

type UpdateUserInfoCond struct {
	Password    *string
	Nickname    *string
	Birthdate   *time.Time
	Gender      *int
	CountryCode *string
	Address     *string
	PhoneNumber *string
}
