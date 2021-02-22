package main

import (
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
)

func catchSig(ch chan os.Signal, c *websocket.Conn) {
	<- ch
	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("write close: ", err)
	}
	return
}

func main() {
	interrupt := make(chan os.Signal,1)
	signal.Notify(interrupt, os.Interrupt)

	u:= "ws://localhost:8000/"
	log.Printf("connecting to %s", u)

	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal("dial:" , err)
	}
	defer c.Close()

	go catchSig(interrupt, c)

	process(c)

}