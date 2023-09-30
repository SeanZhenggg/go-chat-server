package controllers

import (
	"chat/app/constants"
	"chat/app/model/bo"
	"chat/app/model/dto"
	"chat/app/service"
	"chat/app/utils/logger"
	"encoding/json"
	"net/http"
	"time"

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

	// user login validation
	//boUserInfo, err := ctrl.userSrv.ValidateUser(&bo.UserValidateCond{Token: chatQueryDto.Token})
	//if err != nil || chatQueryDto.Account != boUserInfo.Account {
	//	ctrl.logger.Error(xerrors.Errorf("Conn ValidateUser error : %w", err))
	//	return
	//}

	// test concurrent code
	boUserInfo := &bo.UserInfo{
		Id:       0,
		Account:  chatQueryDto.Account,
		Password: "test1234",
		Nickname: chatQueryDto.Account,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	conn, err := ctrl.defaultUpgrade().Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		ctrl.logger.Error(xerrors.Errorf("Conn Websocket Connection error : %w", err))
		return
	}

	roomId := bo.RoomId(chatQueryDto.RoomId)
	client := &bo.ClientState{
		IsRegister: constants.ClientState_Registered,
		Client: &bo.Client{
			UserInfo: boUserInfo,
			Conn:     conn,
			RoomId:   roomId,
		},
	}

	ctrl.hubSrv.GetRoomOrCreateIfNotExisted(roomId)

	ctrl.hubSrv.HouseChange(client)

	chatMessage := bo.ChatMessage{
		RoomId:   bo.RoomId(chatQueryDto.RoomId),
		Account:  chatQueryDto.Account,
		Message:  "已連線",
		Nickname: chatQueryDto.Account,
	}

	message, err := json.Marshal(chatMessage)
	if err != nil {
		ctrl.logger.Error(xerrors.Errorf("JSON marshal error : %w", err))
	}

	ctrl.hubSrv.BroadCastMsg(&bo.BroadcastState{
		Message: message,
		RoomId:  bo.RoomId(chatQueryDto.RoomId),
	})

	go ctrl.readPump(client.Client)
}

func (ctrl *ChatCtrl) defaultUpgrade() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

func (ctrl *ChatCtrl) readPump(cli *bo.Client) {
	defer func() {
		ctrl.hubSrv.HouseChange(&bo.ClientState{
			Client:     cli,
			IsRegister: constants.ClientState_UnRegistered,
		})
		cli.Conn.Close()
	}()

	cli.Conn.SetReadDeadline(time.Now().Add(constants.PongWait))
	cli.Conn.SetPongHandler(func(a string) error {
		cli.Conn.SetReadDeadline(time.Now().Add(constants.PongWait))
		return nil
	})

	for {
		_, message, err := cli.Conn.ReadMessage()
		if err != nil {
			ctrl.logger.Error(xerrors.Errorf("readPump IsUnexpectedCloseError error : %w", err))
			break
		}

		chatMessage := bo.ChatMessage{
			RoomId:   cli.RoomId,
			Account:  cli.UserInfo.Account,
			Message:  string(message),
			Nickname: cli.UserInfo.Nickname,
		}

		sendMsg, err := json.Marshal(chatMessage)
		if err != nil {
			ctrl.logger.Error(xerrors.Errorf("JSON marshal error : %w", err))
		}

		ctrl.hubSrv.BroadCastMsg(&bo.BroadcastState{
			Message: sendMsg,
			RoomId:  cli.RoomId,
		})
	}
}
