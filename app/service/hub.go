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

type hub struct {
	clients       map[*bo.Client]struct{}
	roomId        bo.RoomId
	clientChan    chan *bo.ClientState
	broadcastChan chan *bo.BroadcastState
	mu            sync.Mutex
	logger        logger.ILogger
}

func (srv *hub) run(hubSrv *hubService) {
	ticker := time.NewTicker(constants.PingPeriod)

	defer func() {
		ticker.Stop()
		close(srv.clientChan)
		close(srv.broadcastChan)
		fmt.Println("run stop ...")
	}()

	for {
		select {
		case state := <-srv.clientChan:
			srv.mu.Lock()

			switch state.IsRegister {
			case constants.ClientState_Registered:
				srv.clients[state.Client] = struct{}{}
			case constants.ClientState_UnRegistered:
				delete(srv.clients, state.Client)
			}

			fmt.Printf("room clients count : %v\n", len(srv.clients))

			if len(srv.clients) == 0 {
				hubSrv.house.Delete(state.Client.RoomId)
				srv.mu.Unlock()
				fmt.Println("delete room")
				return
			}

			srv.mu.Unlock()
		case message := <-srv.broadcastChan:
			srv.broadcastMsg(hubSrv, message)
		case <-ticker.C:
			srv.sendPingToAll()
		}
	}
}

func (srv *hubService) createHub(roomId bo.RoomId) *hub {
	newHub := &hub{
		clients:       make(map[*bo.Client]struct{}),
		roomId:        roomId,
		clientChan:    make(chan *bo.ClientState, 1024),
		broadcastChan: make(chan *bo.BroadcastState, 1024),
		logger:        srv.logger,
	}

	go newHub.run(srv)

	return newHub
}

func (srv *hubService) HouseChange(data *bo.ClientState) {
	room := srv.getRoom(data.Client.RoomId)
	room.clientChan <- data
}

func (srv *hubService) BroadCastMsg(data *bo.BroadcastState) {
	room := srv.getRoom(data.RoomId)
	if room == nil {
		return
	}

	room.broadcastChan <- data
}

//func (srv *hub) updateClient(data *bo.ClientState) (removed bool) {
//	srv.mu.Lock()
//	defer func() {
//		srv.mu.Unlock()
//		fmt.Println("updateClient end ...\n")
//	}()
//
//	switch data.IsRegister {
//	case constants.ClientState_Registered:
//		srv.clients[data.Client] = struct{}{}
//	case constants.ClientState_UnRegistered:
//		delete(srv.clients, data.Client)
//		if len(srv.clients) == 0 {
//			removed = true
//		}
//	}
//
//	fmt.Println("🍎🍎🍎🍎🍎🍎 -----clients now start-----")
//	for client := range srv.clients {
//		fmt.Printf("🍎🍎🍎🍎🍎🍎 %v \n", client.UserInfo.Account)
//	}
//	fmt.Println("🍎🍎🍎🍎🍎🍎 -----clients now end-----")
//
//	return
//}

func (srv *hubService) GetRoomOrCreateIfNotExisted(roomId bo.RoomId) *hub {
	room, isExist := srv.house.Load(roomId)
	if !isExist {
		newHub := srv.createHub(roomId)
		srv.house.Store(roomId, newHub)
		room = newHub
	}

	return room.(*hub)
}

func (srv *hubService) getRoom(roomId bo.RoomId) *hub {
	room, isExist := srv.house.Load(roomId)
	if !isExist {
		srv.logger.Error(xerrors.Errorf("getRoom no room found : %v", roomId))
		return nil
	}

	return room.(*hub)
}

func (srv *hub) broadcastMsg(hubSrv *hubService, data *bo.BroadcastState) {
	for client := range srv.clients {
		srv.writeMsg(hubSrv, client, data)
	}
}

func (srv *hub) writeMsg(hubSrv *hubService, client *bo.Client, data *bo.BroadcastState) {
	var err error
	if err = client.Conn.SetWriteDeadline(time.Now().Add(constants.WriteWait)); err != nil {
		srv.logger.Error(xerrors.Errorf("broadcastMsg SetWriteDeadline error : %w", err))
		//hubSrv.closeClientConn(client)
		return
	}

	w, err := client.Conn.NextWriter(websocket.TextMessage)

	defer func() {
		if w != nil {
			w.Close()
		}
	}()

	if err != nil {
		//srv.logger.Error(xerrors.Errorf("writeMsg NextWriter error : %w", err))
		//hubSrv.closeClientConn(client)
		return
	}

	if _, err = w.Write(data.Message); err != nil {
		//srv.logger.Error(xerrors.Errorf("writeMsg Write error : %w", err))
		//hubSrv.closeClientConn(client)
		return
	}

}

func (srv *hub) sendPingToAll() {
	srv.mu.Lock()
	defer srv.mu.Unlock()

	for client := range srv.clients {
		if err := client.Conn.SetWriteDeadline(time.Now().Add(constants.WriteWait)); err != nil {
			srv.logger.Error(xerrors.Errorf("broadcastMsg SetWriteDeadline error : %w", err))
			return
		}

		if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
			return
		}
	}
}

func (srv *hubService) closeClientConn(client *bo.Client) {
	//fmt.Printf("call closeClientConn from client : %v\n", client.UserInfo.Account)
	//srv.HouseChange(&bo.ClientState{
	//	Client:     client,
	//	IsRegister: constants.ClientState_UnRegistered,
	//})
	//client.Conn.Close()
}
