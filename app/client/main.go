package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var (
	addr      = flag.String("addr", "localhost:8080", "http service address")
	interrupt = make(chan os.Signal, 1)
)

func main() {
	for i := 0; i < 500; i++ {
		go createClient(fmt.Sprintf("account=sean00%v&roomId=1&token=aaa", i))
	}

	<-interrupt
}

func createClient(query string) {
	flag.Parse()
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/websocket", RawQuery: query}
	// u := url.URL{Scheme: "ws", Host: *addr, Path: "/websocket"}
	log.Printf("connecting to %s", u.String())

	c, err := connectWS(u.String())
	if err != nil {
		log.Fatal("dial:", err)
		return
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Millisecond * 1000)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func _connection(url string) (*websocket.Conn, error) {
	c, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		log.Printf("connection error : %v", err)
		fmt.Printf("c : %v", c)

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
