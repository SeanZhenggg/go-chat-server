package dto

import "time"

type UserCondDto struct {
	Account string `uri:"account"`
}

type UserDto struct {
	Id       uint      `json:"id"`
	Account  string    `json:"account"`
	Nickname string    `json:"nickname"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type UserRegDto struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type UserLoginDto struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type UserLoginRespDto struct {
	Token string `json:"token"`
}
