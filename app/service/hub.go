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
	GetRoomOrCreateIfNotExisted(roomId bo.RoomId) *hub
	ClientStateChange(data *bo.ClientState)
	BroadCastMsg(data *bo.BroadcastState)
}

func ProvideHubSrv(logger logger.ILogger) IHubSrv {
	obj := &hubService{
		logger:    logger,
		roomLocks: make(map[bo.RoomId]*sync.Mutex),
	}

	return obj
}

type hubService struct {
	house      sync.Map
	roomMuLock sync.Mutex
	roomLocks  map[bo.RoomId]*sync.Mutex
	logger     logger.ILogger
}

func (srv *hubService) createHub(roomId bo.RoomId) *hub {
	roomLock := &sync.Mutex{}
	srv.roomLocks[roomId] = roomLock

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

func (srv *hubService) ClientStateChange(data *bo.ClientState) {
	room := srv.GetRoomOrCreateIfNotExisted(data.RoomId)
	fmt.Printf("ClientStateChange room %+v\n", room)
	room.clientStateChan <- data
}

func (srv *hubService) BroadCastMsg(data *bo.BroadcastState) {
	room, err := srv.getRoom(data.RoomId)
	if err != nil {
		srv.logger.Error(xerrors.Errorf("BroadCastMsg error : %w", err))
		return
	}

	room.broadcastStateChan <- data
}

func (srv *hubService) GetRoomOrCreateIfNotExisted(roomId bo.RoomId) *hub {
	srv.roomMuLock.Lock()
	defer srv.roomMuLock.Unlock()

	var (
		room *hub
		err  error
	)

	room, err = srv.getRoom(roomId)

	fmt.Println("Sleep 5 seconds...")
	time.Sleep(time.Second * 5)
	if err != nil {
		fmt.Printf("Create room %v...\n", roomId)
		room = srv.createHub(roomId)
		srv.house.Store(roomId, room)
	} else {
		fmt.Printf("Found room %v...\n", roomId)
	}

	return room
}

func (srv *hubService) getRoom(roomId bo.RoomId) (*hub, error) {
	room, isExist := srv.house.Load(roomId)

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
	clientStateChan    chan *bo.ClientState
	broadcastStateChan chan *bo.BroadcastState
	hubSrv             *hubService
}

func (srv *hub) run() {
	ticker := time.NewTicker(constants.PingPeriod)

	defer func() {
		ticker.Stop()
		close(srv.clientStateChan)
		close(srv.broadcastStateChan)
	}()

	for {
		select {
		case state := <-srv.clientStateChan:
			roomLock := srv.hubSrv.roomLocks[state.RoomId]
			roomLock.Lock()

			srv.updateClient(state)
			if len(srv.clients) == 0 {
				delete(srv.clients, state.Client)
				srv.hubSrv.house.Delete(state.RoomId)
				fmt.Printf("Delete room %v...\n", state.RoomId)
				roomLock.Unlock()

				srv.hubSrv.roomMuLock.Lock()
				if roomLock.TryLock() {
					if len(srv.clients) == 0 {
						delete(srv.hubSrv.roomLocks, state.RoomId)
					}
					roomLock.Unlock()
				}

				srv.hubSrv.roomMuLock.Unlock()
				return
			}

			roomLock.Unlock()
		case message := <-srv.broadcastStateChan:
			srv.broadcastMsg(message)
		case <-ticker.C:
			srv.sendPingToAll()
		}
	}
}

func (srv *hub) updateClient(data *bo.ClientState) {
	switch data.IsRegister {
	case constants.ClientState_Registered:
		srv.clients[data.Client] = struct{}{}
	case constants.ClientState_UnRegistered:
		delete(srv.clients, data.Client)
	}
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
