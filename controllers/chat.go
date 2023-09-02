package controllers

import (
	"chat/constants"
	"chat/model/bo"
	"chat/model/dto"
	"chat/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type IChatCtrl interface {
	Conn(ctx *gin.Context)
}

func ProvideChatCtrl(hubSrv service.IHubSrv, userSrv service.IUserSrv) IChatCtrl {
	return &ChatCtrl{
		hubSrv:  hubSrv,
		userSrv: userSrv,
	}
}

type ChatCtrl struct {
	hubSrv  service.IHubSrv
	userSrv service.IUserSrv
}

func (ctrl *ChatCtrl) Conn(ctx *gin.Context) {
	// ctrl.userSrv.
	chatQueryDto := &dto.ChatQueryDto{}
	if err := ctx.BindQuery(chatQueryDto); err != nil {
		fmt.Printf("🍎🍎🍎🍎🍎🍎 Conn BindQuery error : %v\n", err)
		return
	}

	conn, err := ctrl.defaultUpgrade().Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		fmt.Printf("🍎🍎🍎🍎🍎🍎 Conn websocket connection error : %v\n", err)
		return
	}

	client := &bo.ClientState{
		IsRegister: constants.ClientState_Registered,
		Client: &bo.Client{
			// UserInfo: ,
			Conn: conn,
		},
	}
	ctrl.hubSrv.ClientChange(client)

	// connectedMsg := &ctrl.hubSrv.BroadCastMsg()
}

func (ctrl *ChatCtrl) defaultUpgrade() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
}

func (ctrl *ChatCtrl) readPump() {

}
