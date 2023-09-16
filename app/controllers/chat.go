package controllers

import (
	"chat/app/constants"
	"chat/app/model/bo"
	"chat/app/model/dto"
	"chat/app/service"
	"chat/app/utils/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/xerrors"
)

type IChatCtrl interface {
	Conn(ctx *gin.Context)
}

func ProvideChatCtrl(hubSrv service.IHubSrv, userSrv service.IUserSrv, logger logger.ILogger) IChatCtrl {
	return &ChatCtrl{
		hubSrv:  hubSrv,
		userSrv: userSrv,
		logger:  logger,
	}
}

type ChatCtrl struct {
	hubSrv  service.IHubSrv
	userSrv service.IUserSrv
	logger  logger.ILogger
}

func (ctrl *ChatCtrl) Conn(ctx *gin.Context) {
	chatQueryDto := &dto.ChatQueryDto{}
	if err := ctx.BindQuery(chatQueryDto); err != nil {
		ctrl.logger.Error(xerrors.Errorf("Conn BindQuery error : %w", err))
		return
	}

	boUserInfo, err := ctrl.userSrv.ValidateUser(&bo.UserValidateCond{Token: chatQueryDto.Token})
	if err != nil || chatQueryDto.Account != boUserInfo.Account {
		ctrl.logger.Error(xerrors.Errorf("Conn ValidateUser error : %w", err))
		return
	}

	conn, err := ctrl.defaultUpgrade().Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		ctrl.logger.Error(xerrors.Errorf("Conn Websocket Connection error : %w", err))
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
				ctrl.logger.Error(xerrors.Errorf("readPump IsUnexpectedCloseError error : %w", err))
			}
			break
		}

		ctrl.hubSrv.BroadCastMsg(&bo.BroadcastState{
			Message: message,
			RoomId:  cli.RoomId,
		})
	}
}
