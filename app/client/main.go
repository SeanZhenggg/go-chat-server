package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

var (
	addr = flag.String("addr", "localhost:8080", "http service address")
)

func main() {
	for i := 0; i < 300; i++ {
		go createClient(fmt.Sprintf("account=sean00%v&roomId=1&token=aaa", i))
		//go createClient("")
	}

	select {}
}

func createClient(query string) {
	flag.Parse()

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/websocket", RawQuery: query}
	//u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws/1", RawQuery: query}
	log.Printf("connecting to %s", u.String())

	c, err := connectWS(u.String())
	if err != nil {
		log.Fatal("dial:", err)
		return
	}

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

func _connection(url string) (*websocket.Conn, error) {
	c, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		log.Printf("connection error : %v", err)

		return nil, err
	}

	return c, nil
}

func connectWS(url string) (*websocket.Conn, error) {
	var c *websocket.Conn
	var err error

	for count := 0; count < 5; count++ {
		c, err = _connection(url)
		if err == nil {
			return c, nil
		}
		time.After(1 * time.Second)
	}

	log.Printf("connection failed after 5 times...")

	return nil, err
}
