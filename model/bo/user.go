package bo

import "time"

type UserCond struct {
	Account string `json:"account"`
}

type UserLoginData struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type UserRegData struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type UserInfo struct {
	Id       uint      `json:"id"`
	Account  string    `json:"account"`
	Password string    `json:"password"`
	Nickname string    `json:"nickname"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type UserLoginResp struct {
	Token string `json:"token"`
}
