package bo

import (
	"github.com/gorilla/websocket"
)

type RoomId string

type Room map[string]map[*Client]struct{}

type Client struct {
	Conn     *websocket.Conn
	UserInfo *UserInfo
	RoomId   RoomId
}

type ClientState struct {
	IsRegister bool
	Client     *Client
}

type RoomState struct {
	IsJoin bool
	Client *Client
	RoomId RoomId
}

type BroadcastState struct {
	Message []byte
	RoomId  RoomId
}
