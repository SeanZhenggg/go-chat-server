package constants

import (
	"time"
)

const (
	ClientState_Registered   bool = true
	ClientState_UnRegistered bool = false
)

const (
	RoomClientState_Join  bool = true
	RoomClientState_Leave bool = false
)

const (
	WriteWait  = 10 * time.Second
	PongWait   = 60 * time.Second
	PingPeriod = (PongWait * 9) / 10
)
