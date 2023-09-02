package dto

type ChatQueryDto struct {
	Account string `form:"account"`
	RoomId  string `form:"roomId"`
	Token   string `form:"token"`
}
