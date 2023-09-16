package service

import (
	"chat/app/constants"
	"chat/app/model/bo"
	"chat/app/utils/logger"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/xerrors"
)

type IHubSrv interface {
	ClientChange(data *bo.ClientState)
	BroadCastMsg(data *bo.BroadcastState)
	GetRoomOrCreateIfNotExisted(roomId bo.RoomId) map[*bo.Client]struct{}
}

func ProvideHubSrv(logger logger.ILogger) IHubSrv {
	obj := &hubService{
		clients:       make(map[*bo.Client]struct{}),
		rooms:         make(map[bo.RoomId]map[*bo.Client]struct{}, 100),
		clientChan:    make(chan *bo.ClientState, 256),
		broadcastChan: make(chan *bo.BroadcastState, 256),
		logger:        logger,
	}

	go obj.run()

	return obj
}

type hubService struct {
	clients       map[*bo.Client]struct{}
	rooms         map[bo.RoomId]map[*bo.Client]struct{}
	clientChan    chan *bo.ClientState
	broadcastChan chan *bo.BroadcastState
	logger        logger.ILogger
}

func (srv *hubService) run() {
	for {
		select {
		case state := <-srv.clientChan:
			srv.updateClient(state)
		case message := <-srv.broadcastChan:
			srv.broadcastMsg(message)
		}
	}
}

func (srv *hubService) ClientChange(data *bo.ClientState) {
	srv.clientChan <- data
}

func (srv *hubService) BroadCastMsg(data *bo.BroadcastState) {
	srv.broadcastChan <- data
}

func (srv *hubService) updateClient(data *bo.ClientState) {
	switch data.IsRegister {
	case constants.ClientState_Registered:
		srv.clients[data.Client] = struct{}{}
	case constants.ClientState_UnRegistered:
		delete(srv.clients, data.Client)
	}

	srv.updateRoom(&bo.RoomState{
		Client: data.Client,
		IsJoin: data.IsRegister,
		RoomId: data.Client.RoomId,
	})
}

func (srv *hubService) GetRoomOrCreateIfNotExisted(roomId bo.RoomId) map[*bo.Client]struct{} {
	return srv.getRoomOrCreateIfNotExisted(roomId)
}

func (srv *hubService) getRoomOrCreateIfNotExisted(roomId bo.RoomId) map[*bo.Client]struct{} {
	room, isExist := srv.rooms[roomId]
	if !isExist {
		srv.rooms[roomId] = make(map[*bo.Client]struct{}, 1024)
		room = srv.rooms[roomId]
	}

	return room
}

func (srv *hubService) updateRoom(data *bo.RoomState) {
	room := srv.getRoomOrCreateIfNotExisted(data.RoomId)

	switch data.IsJoin {
	case constants.RoomClientState_Join:
		room[data.Client] = struct{}{}
	case constants.RoomClientState_Leave:
		delete(room, data.Client)

		// clean empty room
		if len(room) == 0 {
			delete(srv.rooms, data.RoomId)
		}
	}
}

func (srv *hubService) broadcastMsg(data *bo.BroadcastState) {
	room := srv.getRoom(data.RoomId)
	if room == nil {
		return
	}

	for client := range room {
		if err := client.Conn.SetWriteDeadline(time.Now().Add(constants.WriteWait)); err != nil {
			srv.logger.Error(xerrors.Errorf("broadcastMsg SetWriteDeadline error : %w", err))
			continue
		}

		w, err := client.Conn.NextWriter(websocket.TextMessage)
		if err != nil {
			srv.logger.Error(xerrors.Errorf("broadcastMsg NextWriter error : %w", err))
			continue
		}

		if _, err := w.Write(data.Message); err != nil {
			srv.logger.Error(xerrors.Errorf("broadcastMsg Write error : %w", err))
		}

		if err := w.Close(); err != nil {
			srv.logger.Error(xerrors.Errorf("broadcastMsg Close error : %w", err))
		}
	}
}

func (srv *hubService) getRoom(roomId bo.RoomId) map[*bo.Client]struct{} {
	room, ok := srv.rooms[roomId]
	if !ok {
		srv.logger.Error(xerrors.Errorf("getRoom no room found : %v", roomId))
		return nil
	}

	return room
}