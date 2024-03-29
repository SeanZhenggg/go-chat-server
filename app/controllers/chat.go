package controllers

import (
	"chat/app/constants"
	"chat/app/model/bo"
	"chat/app/model/dto"
	"chat/app/service"
	ctxUtil "chat/app/utils/ctx"
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

	boUserInfo, err := ctxUtil.GetUserInfo(ctx)
	if err != nil {
		return
	}

	// test concurrent code
	//boUserInfo := &bo.UserInfo{
	//	Id:       0,
	//	Account:  chatQueryDto.Account,
	//	Password: "test1234",
	//	Nickname: chatQueryDto.Account,
	//	CreateAt: time.Now(),
	//	UpdateAt: time.Now(),
	//}

	conn, err := ctrl.defaultUpgrade().Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		ctrl.logger.Error(xerrors.Errorf("Conn Websocket Connection error : %w", err))
		return
	}

	roomId := bo.RoomId(chatQueryDto.RoomId)
	client := &bo.Client{
		UserInfo: boUserInfo,
		Conn:     conn,
		RoomId:   roomId,
	}
	ctrl.hubSrv.ClientRegister(client)

	chatMessage := &bo.ChatMessage{
		RoomId:   bo.RoomId(chatQueryDto.RoomId),
		Account:  chatQueryDto.Account,
		Message:  "已連線",
		Nickname: chatQueryDto.Account,
	}
	message, err := json.Marshal(chatMessage)
	if err != nil {
		ctrl.logger.Error(xerrors.Errorf("JSON marshal error : %w", err))
	}
	ctrl.hubSrv.BroadcastMsg(&bo.BroadcastState{
		Message: message,
		RoomId:  bo.RoomId(chatQueryDto.RoomId),
	})

	go ctrl.readPump(client)
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
		ctrl.hubSrv.ClientUnregister(cli)
		cli.Conn.Close()
	}()

	cli.Conn.SetReadDeadline(time.Now().Add(constants.PongWait))
	cli.Conn.SetPongHandler(func(a string) error {
		//fmt.Printf("receive pong ... %v", a)
		cli.Conn.SetReadDeadline(time.Now().Add(constants.PongWait))
		return nil
	})

	for {
		_, message, err := cli.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				ctrl.logger.Error(xerrors.Errorf("readPump IsUnexpectedCloseError error : %w", err))
			}
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
			break
		}

		ctrl.hubSrv.BroadcastMsg(&bo.BroadcastState{
			Message: sendMsg,
			RoomId:  cli.RoomId,
		})
	}
}
