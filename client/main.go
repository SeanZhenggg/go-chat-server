package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	mrand "math/rand"
	"net/url"
	"strings"
	"time"
)

var (
	account = flag.String("acc", "sean001", "account to join room")
	roomId  = flag.Int("room", 1, "connected to which room")
	addr    = flag.String("addr", "localhost:8080", "http service address")
)

func main() {
	var accounts []string
	flag.Parse()

	if strings.ContainsAny(*account, ",") {
		accounts = append(accounts, strings.Split(*account, ",")...)
	} else {
		accounts[0] = *account
	}

	fmt.Printf("accounts: %+v, roomId: %d, addr: %s\n", accounts, *roomId, *addr)

	//concurrentClientMsgSendTest(accounts)
	concurrentJoinAndLeavingTest(accounts)

	select {}

}

func concurrentClientMsgSendTest(accs []string) {
	for i := 0; i < 300; i++ {
		randIdx := mrand.Intn(len(accs))
		acc := accs[randIdx]
		t, err := generateRandomToken(64)
		if err != nil {
			fmt.Printf("token error : %v\n", err)
		}

		go func() {
			c, err := createClient(fmt.Sprintf("account=%v&roomId=%v&token=%v", acc, *roomId, t))
			if err != nil {
				fmt.Printf("create client error : %v\n", err)
				return
			}

			sendAndRecvMessage(c)
		}()
	}
}

func concurrentJoinAndLeavingTest(accs []string) {
	// join and leave
	go func() {
		t, err := generateRandomToken(64)
		if err != nil {
			fmt.Printf("token error : %v\n", err)
		}

		joinAndLeave(accs[0], t)
	}()

	// join
	go func() {
		t, err := generateRandomToken(64)
		if err != nil {
			fmt.Printf("token error : %v\n", err)
		}
		//time.Sleep(20 * time.Millisecond)
		//
		//_, err = createClient(fmt.Sprintf("account=%v&roomId=%v&token=%v", accs[1], *roomId, t))
		//if err != nil {
		//	fmt.Printf("create client error : %v\n", err)
		//	return
		//}
		joinAndLeave(accs[1], t)

		//fmt.Printf("=================account %v join room %v =====================\n", accs[1], *roomId)
	}()
}

func createClient(query string) (*websocket.Conn, error) {

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/websocket", RawQuery: query}
	//u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws/1", RawQuery: query}
	log.Printf("connecting to %s", u.String())

	c, err := connectWS(u.String())

	return c, err
}

func sendAndRecvMessage(c *websocket.Conn) {
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

func joinAndLeave(acc string, t string) {
	var c *websocket.Conn = nil
	var err error = nil
	for {
		c, err = createClient(fmt.Sprintf("account=%v&roomId=%v&token=%v", acc, *roomId, t))
		fmt.Printf("account %v joinned room %v\n", acc, *roomId)
		time.Sleep(20 * time.Millisecond)
		if err != nil {
			fmt.Printf("token error : %v\n", err)
		}

		c.Close()
		fmt.Printf("\"=================account %v leaved room %v\"=================\n", acc, *roomId)
		time.Sleep(50 * time.Millisecond)
	}
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
