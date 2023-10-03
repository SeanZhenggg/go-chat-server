package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

var (
	account = flag.String("acc", "sean001", "account to join room")
	roomId  = flag.Int("room", 1, "connected to which room")
	addr    = flag.String("addr", "localhost:8080", "http service address")
)

func main() {
	flag.Parse()
	//fmt.Printf("account: %s, roomId: %d, addr: %s\n", *account, *roomId, *addr)

	for i := 0; i < 300; i++ {
		t, err := generateRandomToken(64)
		if err != nil {
			fmt.Printf("token error : %v\n", err)
		}

		go createClient(fmt.Sprintf("account=%v&roomId=%v&token=%v", *account, *roomId, t))
	}

	select {}

}

func createClient(query string) {

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/websocket", RawQuery: query}
	//u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws/1", RawQuery: query}
	log.Printf("connecting to %s", u.String())

	c, err := connectWS(u.String())
	if err != nil {
		log.Fatal("dial:", err)
		return
	}

	// read
	go func() {
		for {
			select {
			default:
				_, message, err := c.ReadMessage()
				if err != nil {
					log.Println("ReadMessage error:", err)
					return
				}
				log.Printf("recv: %s", message)
			}
		}
	}()

	// write
	go func() {
		ticker := time.NewTicker(time.Millisecond * 1000)
		defer ticker.Stop()

		for {
			select {
			case t := <-ticker.C:
				err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
				if err != nil {
					log.Println("write:", err)
					return
				}
			}
		}
	}()
}

func connectWS(url string) (c *websocket.Conn, err error) {
	for count := 0; count < 5; count++ {
		c, _, err := websocket.DefaultDialer.Dial(url, nil)

		if err != nil {
			log.Printf("connection error : %v", err)
			time.After(1 * time.Second)
			continue
		}

		return c, nil
	}

	log.Printf("connection failed after 5 times...")

	return nil, err
}

func generateRandomToken(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(randomBytes)
	return token, nil
}
