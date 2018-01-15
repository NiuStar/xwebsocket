package xwebsocket

import (
	"golang.org/x/net/websocket"
	"fmt"
)
type Connet func(con *WebSocketContext)
type ReadMsg func(con *WebSocketContext)
type Disconncet func(con *WebSocketContext)
type WebSocketDelegate struct {
	connect Connet
	readMsg ReadMsg
	disconnect Disconncet
}

func NewWebSocketDelagte(c Connet,r ReadMsg , d Disconncet) *WebSocketDelegate {
	delegate := new (WebSocketDelegate)
	delegate.connect = c
	delegate.readMsg = r
	delegate.disconnect = d
	return delegate

}

type WebSocketContext struct {
	Conn *websocket.Conn
	Message string
}

func NewWebSocketContext(c *websocket.Conn) *WebSocketContext {
	delegate := new (WebSocketContext)
	delegate.Conn = c
	return delegate

}

func (c *WebSocketContext)WriteMsg (msg string) bool {//发送给用户con信息msg

	if len(msg) > 100 {
		fmt.Println("send:::",msg[0:100])
	} else {
		fmt.Println("send:::",msg)
	}

	if err := websocket.Message.Send(c.Conn, msg); err != nil {
		return false
	}
	return true
}