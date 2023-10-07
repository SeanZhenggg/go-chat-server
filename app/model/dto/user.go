package dto

import "time"

type UserCondDto struct {
	Account string `uri:"account"`
}

type UserDto struct {
	Id          uint      `json:"id"`
	Account     string    `json:"account"`
	Nickname    string    `json:"nickname"`
	Birthdate   time.Time `json:"birthdate"`
	Gender      int       `json:"gender"`
	Country     string    `json:"country"`
	Address     string    `json:"address"`
	RegionCode  string    `json:"region_code"`
	PhoneNumber string    `json:"phone_number"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
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
