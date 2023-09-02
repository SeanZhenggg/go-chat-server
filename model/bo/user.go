package bo

import "time"

type UserCond struct {
	Account string
}

type UserLoginData struct {
	Account  string
	Password string
}

type UserRegData struct {
	Account  string
	Password string
	Nickname string
}

type UserInfo struct {
	Id       uint
	Account  string
	Password string
	Nickname string
	CreateAt time.Time
	UpdateAt time.Time
}

type UserLoginResp struct {
	Token string
}

type UserValidateCond struct {
	Token string
}
