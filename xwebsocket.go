package xwebsocket

import (
	"fmt"
	"golang.org/x/net/websocket"
)

type XWebSocket struct {
	delegate *WebSocketDelegate
}

func NewXWebSocket(c Connet,r ReadMsg , d Disconncet) *XWebSocket {
	web := new (XWebSocket)
	web.delegate = NewWebSocketDelagte(c,r,d)
	return web

}

// Echo the data received on the WebSocket.
func (this *XWebSocket)EchoServer(ws *websocket.Conn) {

	c := NewWebSocketContext(ws)

	this.delegate.connect(c)
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			this.delegate.disconnect(c)
			break
		}
		c.Message = reply
		this.delegate.readMsg(c)
	}
}

