package controllers

import (
	"chat/app/constants"
	"chat/app/model/bo"
	"chat/app/model/dto"
	"chat/app/service"
	"fmt"

	// "time"

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
	chatQueryDto := &dto.ChatQueryDto{}
	if err := ctx.BindQuery(chatQueryDto); err != nil {
		fmt.Printf("🍎🍎🍎🍎🍎🍎 Conn BindQuery error : %v\n", err)
		return
	}

	boUserInfo, err := ctrl.userSrv.ValidateUser(&bo.UserValidateCond{chatQueryDto.Token})
	if err != nil || boUserInfo.Account != chatQueryDto.Account {
		fmt.Printf("🍎🍎🍎🍎🍎🍎 Conn ValidateUser error : %v\n", err)
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
			UserInfo: boUserInfo,
			Conn:     conn,
			RoomId:   bo.RoomId(chatQueryDto.RoomId),
		},
	}
	ctrl.hubSrv.ClientChange(client)

	ctrl.hubSrv.GetRoomOrCreateIfNotExisted(bo.RoomId(chatQueryDto.RoomId))

	ctrl.hubSrv.BroadCastMsg(&bo.BroadcastState{
		Message: []byte("已連線"),
		RoomId:  bo.RoomId(chatQueryDto.RoomId),
	})

	go ctrl.readPump(client.Client)
}

func (ctrl *ChatCtrl) defaultUpgrade() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
}

func (ctrl *ChatCtrl) readPump(cli *bo.Client) {
	defer func() {
		cli.Conn.Close()
	}()

	// cli.Conn.SetReadDeadline(time.Now().Add(constants.PongWait))
	// cli.Conn.SetPongHandler(func(string) error { cli.Conn.SetReadDeadline(time.Now().Add(constants.PongWait)); return nil })
	for {
		_, message, err := cli.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("🍎🍎🍎🍎🍎🍎 readPump IsUnexpectedCloseError error : %v\n", err)
			}
			break
		}

		ctrl.hubSrv.BroadCastMsg(&bo.BroadcastState{
			Message: message,
			RoomId:  cli.RoomId,
		})
	}
}
