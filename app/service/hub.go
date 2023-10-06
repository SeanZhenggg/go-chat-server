package service

import (
	"chat/app/constants"
	"chat/app/model/bo"
	"chat/app/utils/logger"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/xerrors"
)

type IHubSrv interface {
	ClientRegister(data *bo.Client)
	ClientUnregister(data *bo.Client)
	BroadcastMsg(data *bo.BroadcastState)
}

func ProvideHubSrv(logger logger.ILogger) IHubSrv {
	obj := &hubService{
		logger:    logger,
		roomLocks: make(map[bo.RoomId]*sync.Mutex),
	}

	return obj
}

type hubService struct {
	rooms            sync.Map
	lockForRoomLocks sync.Mutex
	roomLocks        map[bo.RoomId]*sync.Mutex
	logger           logger.ILogger
}

func (srv *hubService) createHub(roomId bo.RoomId) *hub {
	newHub := &hub{
		clients:            make(map[*bo.Client]struct{}),
		roomId:             roomId,
		unregisterChan:     make(chan *bo.Client, 256),
		broadcastStateChan: make(chan *bo.BroadcastState, 256),
		hubSrv:             srv,
	}

	go newHub.run()

	return newHub
}

func (srv *hubService) ClientRegister(data *bo.Client) {
	srv.lockForRoomLocks.Lock()

	var (
		room   *hub
		err    error
		roomId = data.RoomId
	)

	roomLock, ok := srv.roomLocks[roomId]
	if !ok {
		roomLock = &sync.Mutex{}
		srv.roomLocks[roomId] = roomLock
	}
	srv.lockForRoomLocks.Unlock()
	roomLock.Lock()

	room, err = srv.getRoom(roomId)
	fmt.Println("Sleep 5 seconds...")
	time.Sleep(time.Second * 5)
	if err != nil {
		fmt.Printf("Create room %v...\n", roomId)
		room = srv.createHub(roomId)
		srv.rooms.Store(roomId, room)
	} else {
		fmt.Printf("Found room %v...\n", roomId)
	}

	fmt.Printf("ClientRegister room %+v, mem: %p\n", room, room)
	room.clients[data] = struct{}{}
	roomLock.Unlock()
}

func (srv *hubService) ClientUnregister(data *bo.Client) {
	room, err := srv.getRoom(data.RoomId)
	if err != nil {
		srv.logger.Error(xerrors.Errorf("ClientUnregister error : %w", err))
		return
	}
	room.unregisterChan <- data
}

func (srv *hubService) BroadcastMsg(data *bo.BroadcastState) {
	room, err := srv.getRoom(data.RoomId)
	if err != nil {
		srv.logger.Error(xerrors.Errorf("BroadcastMsg error : %w", err))
		return
	}

	room.broadcastStateChan <- data
}

func (srv *hubService) getRoom(roomId bo.RoomId) (*hub, error) {
	room, isExist := srv.rooms.Load(roomId)

	if !isExist {
		err := xerrors.Errorf("getRoom no room found : %v", roomId)
		srv.logger.Error(err)
		return nil, err
	}

	return room.(*hub), nil
}

type hub struct {
	clients            map[*bo.Client]struct{}
	roomId             bo.RoomId
	unregisterChan     chan *bo.Client
	broadcastStateChan chan *bo.BroadcastState
	hubSrv             *hubService
}

func (srv *hub) run() {
	ticker := time.NewTicker(constants.PingPeriod)

	defer func() {
		ticker.Stop()
		close(srv.unregisterChan)
		close(srv.broadcastStateChan)
	}()

	for {
		select {
		case state := <-srv.unregisterChan:
			srv.unregister(state)
		case message := <-srv.broadcastStateChan:
			srv.broadcast(message)
		case <-ticker.C:
			srv.sendPingToAll()
		}
	}
}

func (srv *hub) unregister(client *bo.Client) {
	roomLock := srv.hubSrv.roomLocks[client.RoomId]
	roomLock.Lock()

	delete(srv.clients, client)

	chatMessage := &bo.ChatMessage{
		RoomId:   client.RoomId,
		Account:  client.UserInfo.Account,
		Message:  "已離線",
		Nickname: client.UserInfo.Nickname,
	}
	message, err := json.Marshal(chatMessage)
	if err != nil {
		srv.hubSrv.logger.Error(xerrors.Errorf("JSON marshal error : %w", err))
	}
	srv.broadcast(&bo.BroadcastState{
		Message: message,
		RoomId:  client.RoomId,
	})

	if len(srv.clients) == 0 {
		srv.hubSrv.rooms.Delete(client.RoomId)
		fmt.Printf("Delete room %v...\n", client.RoomId)
		roomLock.Unlock()

		srv.hubSrv.lockForRoomLocks.Lock()
		if roomLock.TryLock() {
			if len(srv.clients) == 0 {
				delete(srv.hubSrv.roomLocks, client.RoomId)
			}
			roomLock.Unlock()
		}

		srv.hubSrv.lockForRoomLocks.Unlock()
		return
	}

	roomLock.Unlock()
}

func (srv *hub) broadcast(data *bo.BroadcastState) {
	for client := range srv.clients {
		srv.writeMsg(client, data)
	}
}

func (srv *hub) writeMsg(client *bo.Client, data *bo.BroadcastState) {
	var err error
	if err = client.Conn.SetWriteDeadline(time.Now().Add(constants.WriteWait)); err != nil {
		srv.hubSrv.logger.Error(xerrors.Errorf("writeMsg SetWriteDeadline error : %w", err))
		return
	}

	w, err := client.Conn.NextWriter(websocket.TextMessage)

	defer func() {
		if w != nil {
			w.Close()
		}
	}()

	if err != nil {
		//srv.hubSrv.logger.Error(xerrors.Errorf("writeMsg NextWriter error : %w", err))
		return
	}

	if _, err = w.Write(data.Message); err != nil {
		srv.hubSrv.logger.Error(xerrors.Errorf("writeMsg Write error : %w", err))
	}
}

func (srv *hub) sendPingToAll() {
	for client := range srv.clients {
		if err := client.Conn.SetWriteDeadline(time.Now().Add(constants.WriteWait)); err != nil {
			srv.hubSrv.logger.Error(xerrors.Errorf("sendPingToAll SetWriteDeadline error : %w", err))
			return
		}

		if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
			return
		}
	}
}
