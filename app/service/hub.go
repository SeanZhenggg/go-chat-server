package service

import (
	"chat/app/constants"
	"chat/app/model/bo"
	"chat/app/utils/logger"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/xerrors"
)

type IHubSrv interface {
	GetRoomOrCreateIfNotExisted(roomId bo.RoomId)
	HouseChange(data *bo.ClientState)
	BroadCastMsg(data *bo.BroadcastState)
}

func ProvideHubSrv(logger logger.ILogger) IHubSrv {
	obj := &hubService{
		logger: logger,
	}

	return obj
}

type hubService struct {
	house  sync.Map
	logger logger.ILogger
}

func (srv *hubService) createHub(roomId bo.RoomId) *hub {
	newHub := &hub{
		clients:            make(map[*bo.Client]struct{}),
		roomId:             roomId,
		clientStateChan:    make(chan *bo.ClientState, 256),
		broadcastStateChan: make(chan *bo.BroadcastState, 256),
		hubSrv:             srv,
	}

	go newHub.run()

	return newHub
}

func (srv *hubService) HouseChange(data *bo.ClientState) {
	room := srv.getRoom(data.RoomId)
	room.clientStateChan <- data
}

func (srv *hubService) BroadCastMsg(data *bo.BroadcastState) {
	room := srv.getRoom(data.RoomId)
	if room == nil {
		return
	}

	room.broadcastStateChan <- data
}

func (srv *hubService) GetRoomOrCreateIfNotExisted(roomId bo.RoomId) {
	_, isExist := srv.house.Load(roomId)
	if !isExist {
		newHub := srv.createHub(roomId)
		srv.house.Store(roomId, newHub)
	}
}

func (srv *hubService) getRoom(roomId bo.RoomId) *hub {
	room, isExist := srv.house.Load(roomId)
	if !isExist {
		srv.logger.Error(xerrors.Errorf("getRoom no room found : %v", roomId))
		return nil
	}

	return room.(*hub)
}

type hub struct {
	clients            map[*bo.Client]struct{}
	roomId             bo.RoomId
	clientStateChan    chan *bo.ClientState
	broadcastStateChan chan *bo.BroadcastState
	mu                 sync.Mutex
	hubSrv             *hubService
}

func (srv *hub) run() {
	ticker := time.NewTicker(constants.PingPeriod)

	defer func() {
		ticker.Stop()
		close(srv.clientStateChan)
		close(srv.broadcastStateChan)
		fmt.Println("hub closed...")
	}()

	for {
		select {
		case state := <-srv.clientStateChan:
			srv.mu.Lock()

			roomRemoved := srv.updateClient(state)
			if roomRemoved {
				srv.hubSrv.house.Delete(state.RoomId)
				srv.mu.Unlock()
				return
			}

			srv.mu.Unlock()
		case message := <-srv.broadcastStateChan:
			srv.broadcastMsg(message)
		case <-ticker.C:
			srv.sendPingToAll()
		}
	}
}

func (srv *hub) updateClient(data *bo.ClientState) bool {
	removed := false

	switch data.IsRegister {
	case constants.ClientState_Registered:
		srv.clients[data.Client] = struct{}{}
	case constants.ClientState_UnRegistered:
		delete(srv.clients, data.Client)
		if len(srv.clients) == 0 {
			removed = true
		}
	}

	return removed
}

func (srv *hub) broadcastMsg(data *bo.BroadcastState) {
	for client := range srv.clients {
		srv.writeMsg(client, data)
	}
}

func (srv *hub) writeMsg(client *bo.Client, data *bo.BroadcastState) {
	var err error
	if err = client.Conn.SetWriteDeadline(time.Now().Add(constants.WriteWait)); err != nil {
		srv.hubSrv.logger.Error(xerrors.Errorf("broadcastMsg SetWriteDeadline error : %w", err))
		return
	}

	w, err := client.Conn.NextWriter(websocket.TextMessage)

	defer func() {
		if w != nil {
			w.Close()
		}
	}()

	if err != nil {
		srv.hubSrv.logger.Error(xerrors.Errorf("writeMsg NextWriter error : %w", err))
		return
	}

	if _, err = w.Write(data.Message); err != nil {
		srv.hubSrv.logger.Error(xerrors.Errorf("writeMsg Write error : %w", err))
	}
}

func (srv *hub) sendPingToAll() {
	srv.mu.Lock()
	defer srv.mu.Unlock()

	for client := range srv.clients {
		if err := client.Conn.SetWriteDeadline(time.Now().Add(constants.WriteWait)); err != nil {
			srv.hubSrv.logger.Error(xerrors.Errorf("broadcastMsg SetWriteDeadline error : %w", err))
			return
		}

		if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
			return
		}
	}
}
